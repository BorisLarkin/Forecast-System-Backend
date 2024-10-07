package repository

import (
	"web/internal/app/ds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetForecastByID(id int) (*ds.Forecasts, error) {
	forecast := &ds.Forecasts{}

	err := r.db.First(forecast, "id = ?", "1").Error // find product with id = 1
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

func (r *Repository) CreateForecast(forecast ds.Forecasts) error {
	return r.db.Create(forecast).Error
}

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

func (r *Repository) CreateUser(user ds.Users) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByID(id int) (*ds.Users, error) {
	user := &ds.Users{}

	err := r.db.First(user, "id = ?", "1").Error // find product with id = 1
	if err != nil {
		return nil, err
	}

	return user, nil
}
