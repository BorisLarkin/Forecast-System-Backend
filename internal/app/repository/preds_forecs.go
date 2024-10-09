package repository

import (
	"strconv"
	"web/internal/app/ds"
)

func (r *Repository) Preds_forecsList() (*[]ds.Preds_Forecs, error) {
	var Preds_Forecs []ds.Preds_Forecs
	r.db.Find(&Preds_Forecs)
	return &Preds_Forecs, nil
}

func (r *Repository) GetPredForecByID(id1 string, id2 string) (*ds.Preds_Forecs, error) {
	var Preds_Forecs ds.Preds_Forecs
	intId1, _ := strconv.Atoi(id1)
	intId2, _ := strconv.Atoi(id2)
	r.db.Find(&Preds_Forecs, intId1, intId2)
	return &Preds_Forecs, nil
}

func (r *Repository) CreatePreds_Forecs(Preds_Forecs ds.Preds_Forecs) error {
	return r.db.Create(Preds_Forecs).Error
}

func (r *Repository) DeletePreds_Forecs(prediction_id string, forecast_id string) {
	query := "DELETE FROM forecasts WHERE prediction_id = $1 and forecast_id = $2"
	r.db.Exec(query, prediction_id, forecast_id)
}
