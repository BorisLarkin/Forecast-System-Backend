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

func (r *Repository) CreatePreds_Forecs(prediction_id string, forecast_id string) error {
	var n ds.Preds_Forecs
	n.ForecastID, _ = strconv.Atoi(forecast_id)
	n.PredictionID, _ = strconv.Atoi(prediction_id)
	return r.db.Create(n).Error
}

func (r *Repository) DeletePreds_Forecs(prediction_id string, forecast_id string) {
	query := "DELETE FROM forecasts WHERE prediction_id = $1 and forecast_id = $2"
	r.db.Exec(query, prediction_id, forecast_id)
}

func (r *Repository) GetForecastsByID(pred_id string) (*[]ds.Forecasts, error) {
	var prf []ds.Preds_Forecs
	r.db.Where("prediction_id = ?", pred_id).Find(&prf)
	var forecs []ds.Forecasts
	for i := range prf {
		f, err := r.GetForecastByID(strconv.Itoa(prf[i].ForecastID))
		if err != nil {
			return nil, err
		}
		forecs = append(forecs, *f)
	}
	return &forecs, nil
}
func (r *Repository) GetPredLen(pred_id string) int {
	var prf []ds.Preds_Forecs
	r.db.Where("prediction_id = ?", pred_id).Find(&prf)
	return len(prf)
}
