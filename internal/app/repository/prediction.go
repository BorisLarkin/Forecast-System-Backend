package repository

import (
	"strconv"
	"web/internal/app/ds"
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
	if err := r.db.Where("user_id = ? and status= ?", user_id, "draft").First(&Predictions).Error; err != nil {
		return "", err
	}
	aid := strconv.Itoa(Predictions.Id)
	return aid, nil
}

func (r *Repository) CreateDraft(Prediction ds.Predictions) error {
	return r.db.Create(Prediction).Error
}
