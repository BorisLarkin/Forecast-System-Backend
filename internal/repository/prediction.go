package repository

import (
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
	query := "UPDATE predictions SET status=$1 WHERE prediction_id = $2"
	return r.db.Exec(query, status, id).Error
}

func (r *Repository) CreatePrediction(Prediction_ptr *ds.Predictions) error {
	return r.db.Create(Prediction_ptr).Error
}

func (r *Repository) DeletePrediction(prediction_id string) error {
	query := "DELETE FROM predictions WHERE prediction_id = $1"
	return r.db.Exec(query, prediction_id).Error
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
