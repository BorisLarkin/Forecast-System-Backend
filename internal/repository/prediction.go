package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func (r *Repository) PredictionList() (*[]ds.Predictions, error) {
	var Predictions []ds.Predictions
	r.db.Find(&Predictions)
	return &Predictions, nil
}

func (r *Repository) GetPredictionByID(id string) (*ds.Predictions, error) {
	var Prediction ds.Predictions
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Prediction, intId)
	return &Prediction, nil
}

func (r *Repository) SetPredictionStatus(prediction_id string, status string) error {
	var prediction ds.Predictions
	if err := r.db.Model(&ds.Predictions{}).Where("prediction_id = ?", prediction_id).First(&prediction).Error; err != nil {
		return err
	}

	prediction.Status = status
	return r.db.Save(&prediction).Error
}

func (r *Repository) CreatePrediction(Prediction_ptr *ds.Predictions) error {
	return r.db.Create(Prediction_ptr).Error
}

func (r *Repository) FinishPrediction(prediction_id string) error {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, prediction_id).Error; err != nil {
		return err
	}

	prediction.Date_completed = time.Now()
	return r.db.Save(prediction).Error
}

func (r *Repository) DeletePrediction(prediction_id string, creator_id string) error {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, prediction_id).Error; err != nil {
		return err
	}
	int_creator, err := strconv.Atoi(creator_id)
	if err != nil {
		return err
	}
	if prediction.CreatorID != int_creator {
		return errors.New("attempt to change unowned prediction")
	}
	prediction.Status = "deleted"

	return r.db.Save(&prediction).Error
}

func (r *Repository) GetUserDraftID(user_id string) (string, error) {
	var Predictions ds.Predictions
	int_uid, er := strconv.Atoi(user_id)
	if er != nil {
		return "", er
	}
	err := r.db.Where("creator_id=? AND status=?", int_uid, "draft").First(&Predictions).Error
	if err != nil {
		return "", err
	}
	aid := strconv.Itoa(int(Predictions.Prediction_id))
	return aid, nil
}

func (r *Repository) CreateDraft(uid string) error {
	var Prediction ds.Predictions
	intid, _ := strconv.Atoi(uid)
	err := r.db.Where("creator_id=? AND status=?", intid, "draft").First(&Prediction).Error
	if err == nil {
		return fmt.Errorf("draft exists")
	}
	pr := ds.Predictions{CreatorID: intid, Date_created: time.Now(), Status: "draft"}
	return r.CreatePrediction(&pr)
}

func (r *Repository) SavePrediction(id string, ctx *gin.Context) error {
	var Prediction ds.Predictions
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	r.db.Find(&Prediction, int_id)
	Prediction.Date_formed = time.Now()
	Prediction.Prediction_amount, _ = strconv.Atoi(ctx.PostForm("amount"))
	Prediction.Prediction_window, _ = strconv.Atoi(ctx.PostForm("window"))
	val := ctx.PostFormArray("values")
	ids := ctx.PostFormArray("ids")
	r.SaveInputs(int_id, ids, val)
	return r.db.Save(&Prediction).Error
}
func (r *Repository) GetPredictions(uid string, role ds.Role, status string, hasStartDate, hasEndDate bool, startDate, endDate time.Time) (*[]ds.Predictions, error) {
	var predictions []ds.Predictions

	query := r.db.Model(&ds.Predictions{}).Select("predictions.prediction_id, predictions.status, predictions.prediction_amount, predictions.prediction_window, predictions.date_created, predictions.date_formed, predictions.date_completed, predictions.creator_id, predictions.qr")
	if role != ds.Moderator {
		query = query.Where("predictions.creator_id = ?", uid)
	}
	if status != "" {
		query = query.Where("predictions.status = ?", status)
	}
	if hasStartDate {
		query = query.Where("predictions.date_formed >= ?", startDate)
	}
	if hasEndDate {
		query = query.Where("predictions.date_formed <= ?", endDate)
	}

	if err := query.Find(&predictions).Error; err != nil {
		return nil, err
	}

	return &predictions, nil
}

func (r *Repository) EditPrediction(id string, Window int, Amount int) (*ds.Predictions, error) {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, id).Error; err != nil {
		return nil, err
	}

	prediction.Prediction_amount = Amount
	prediction.Prediction_window = Window
	err := r.db.Save(&prediction).Error
	return &prediction, err
}

func (r *Repository) FormPrediction(pred_id string, creatorID string) error {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, pred_id).Error; err != nil {
		return err
	}
	int_creator, err := strconv.Atoi(creatorID)
	if err != nil {
		return err
	}
	if prediction.CreatorID != int_creator {
		return errors.New("attempt to change unowned prediction")
	}

	if prediction.Status != "draft" {
		return errors.New("pre-existing status error")
	}
	prediction.Status = "pending"
	prediction.Date_formed = time.Now()

	return r.db.Save(&prediction).Error
}
func (r *Repository) CalculatePrediction(pred_id string) (*[]ds.Preds_Forecs, error) {
	prediction, err := r.GetPredictionByID(pred_id)
	if err != nil {
		return nil, fmt.Errorf("cannot get prediction")
	}
	preds_forecs, err := r.Preds_forecsList(pred_id)
	if err != nil {
		return nil, fmt.Errorf("cannot get list of predictions_forecasts")
	}
	//process each record
	for i, v := range *preds_forecs {
		output_string := ""
		int_array, err := Calculate(prediction.Prediction_window, prediction.Prediction_amount, v.Input, r.logger)
		if err != nil {
			return nil, fmt.Errorf("issue predicting record (%s, %d): %s", pred_id, i, err)
		}
		for i := range int_array { //stringify each result
			curr_int := strconv.FormatFloat(int_array[i], 'f', 2, 64)
			output_string += curr_int + ","
		}
		output_string = output_string[:len(output_string)-1] //trim the last comma
		v.Result = output_string
		r.db.Save(&v)
	}
	preds_forecs, _ = r.Preds_forecsList(pred_id)
	return preds_forecs, nil
}

func (r *Repository) SetPredictionQr(prediction_id string) error {
	var prediction ds.Predictions
	if err := r.db.Model(&ds.Predictions{}).Where("prediction_id = ?", prediction_id).First(&prediction).Error; err != nil {
		return err
	}
	var preds_forecs, _ = r.Preds_forecsList(prediction_id)

	var png []byte
	var result string
	var creator, _ = r.GetUserByID(strconv.Itoa(prediction.CreatorID))
	result += "Пользователь: " + creator.Login + ", "
	result += "Время отправки: " + prediction.Date_created.Format("2006-01-02 15:04:05") + ", "
	result += "Время завершения: " + prediction.Date_completed.Format("2006-01-02 15:04:05") + ", "

	for _, v := range *preds_forecs {
		f, _ := r.GetForecastByID(strconv.Itoa(int(v.ForecastID)))
		result += "Прогноз: " + f.Short + ", "
		result += "Ввод: " + v.Input + ", "
		result += "Предсказание: " + v.Result + "; "
	}

	png, _ = qrcode.Encode(result, qrcode.Medium, 256)
	prediction.Qr = png
	return r.db.Save(&prediction).Error
}
