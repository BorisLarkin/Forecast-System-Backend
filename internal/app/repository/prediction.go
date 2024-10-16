package repository

import (
	"fmt"
	"strconv"
	"time"
	"web/internal/app/ds"
	"web/internal/app/dsn"
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
	var Prediction ds.Predictions
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Prediction, intId)
	Prediction.Status = status
	r.db.Save(&Prediction)
	return nil
}

func (r *Repository) CreatePrediction(Prediction_ptr *ds.Predictions) error {
	return r.db.Create(Prediction_ptr).Error
}

func (r *Repository) DeletePrediction(prediction_id string) {
	query := "DELETE FROM predictions WHERE id = $1"
	r.db.Exec(query, prediction_id)
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
	aid := strconv.Itoa(Predictions.Id)
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
