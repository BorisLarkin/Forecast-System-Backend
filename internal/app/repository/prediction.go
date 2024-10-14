package repository

import (
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

func (r *Repository) CreatePrediction(Prediction ds.Predictions) error {
	return r.db.Create(Prediction).Error
}

func (r *Repository) DeletePrediction(id string) {
	query := "UPDATE forecasts SET status='deleted' WHERE id = $1"
	r.db.Exec(query, id)
}

func (r *Repository) GetUserDraftID(user_id string) (string, error) {
	var Predictions ds.Predictions
	int_uid, er := strconv.Atoi(user_id)
	if er != nil {
		return "", er
	}
	if err := r.db.Where("user_id=? AND status=?", int_uid, "draft").First(&Predictions).Error; err != nil {
		return "", err
	}
	aid := strconv.Itoa(Predictions.Id)
	return aid, nil
}

func (r *Repository) CreateDraft() error {
	var n ds.Predictions
	uid, _ := dsn.GetCurrentUserID()
	n.UserID, _ = strconv.Atoi(uid)
	n.Date_created = time.Now()
	n.Status = "draft"

	return r.db.Create(n).Error
}
