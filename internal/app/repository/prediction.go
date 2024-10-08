package repository

import "web/internal/app/ds"

func (r *Repository) GetPredictionByID(id int) (*ds.Predictions, error) {
	pred := &ds.Predictions{}

	err := r.db.First(pred, "id = ?", "1").Error // find product with id = 1
	if err != nil {
		return nil, err
	}

	return pred, nil
}

func (r *Repository) CreatePrediction(pred ds.Predictions) error {
	return r.db.Create(pred).Error
}
