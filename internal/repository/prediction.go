package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"web/internal/ds"
	"web/internal/dsn"

	"github.com/gin-gonic/gin"
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

func (r *Repository) SetPredictionStatus(id string, status string) error {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, id).Error; err != nil {
		return err
	}

	prediction.Status = status
	return r.db.Save(&prediction).Error
}

func (r *Repository) CreatePrediction(Prediction_ptr *ds.Predictions) error {
	return r.db.Create(Prediction_ptr).Error
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
	if prediction.UserID != int_creator {
		return errors.New("attempt to change unowned prediction")
	}
	r.SetPredictionStatus(prediction_id, "pending")
	prediction.Date_formed = time.Now()

	return r.db.Save(&prediction).Error
}

func (r *Repository) GetUserDraftID(user_id string) (string, error) {
	var Predictions ds.Predictions
	int_uid, er := strconv.Atoi(user_id)
	if er != nil {
		return "", er
	}
	err := r.db.Where("user_id=? AND status=?", int_uid, "draft").First(&Predictions).Error
	if err != nil {
		return "", err
	}
	aid := strconv.Itoa(Predictions.Prediction_id)
	return aid, nil
}

func (r *Repository) CreateDraft() error {
	var Prediction ds.Predictions
	uid, _ := dsn.GetCurrentUserID()
	intid, _ := strconv.Atoi(uid)
	err := r.db.Where("user_id=? AND status=?", intid, "draft").First(&Prediction).Error
	if err == nil {
		return fmt.Errorf("draft exists")
	}
	pr := ds.Predictions{UserID: intid, Date_created: time.Now(), Status: "draft"}
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
func (r *Repository) GetPredictions(status string, hasStartDate, hasEndDate bool, startDate, endDate time.Time) ([]ds.Predictions, error) {
	var predictions []ds.Predictions

	query := r.db.Table("predictions").
		Select("predictions.prediction_id, predictions.status, predictions.text, predictions.date_created, predictions.date_formed, predictions.date_completed, u1.login as creator").
		Joins("JOIN users u1 ON predictions.user_id = u1.user_id").
		Where("predictions.status != ? AND predictions.status != ?", "deleted", "draft")

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

	return predictions, nil
}

func (r *Repository) EditPrediction(id string, Window int, Amount int) error {
	var prediction ds.Predictions

	if err := r.db.First(&prediction, id).Error; err != nil {
		return err
	}

	prediction.Prediction_amount = Amount
	prediction.Prediction_window = Window
	prediction.Date_formed = time.Now()

	return r.db.Save(&prediction).Error
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
	if prediction.UserID != int_creator {
		return errors.New("attempt to change unowned prediction")
	}

	if prediction.Status != "draft" {
		return errors.New("pre-existing status error")
	}
	r.SetPredictionStatus(pred_id, "pending")
	prediction.Date_formed = time.Now()

	return r.db.Save(&prediction).Error
}
