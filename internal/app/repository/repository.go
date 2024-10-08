package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func New(dsn string, log *logrus.Logger) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:     db,
		logger: log,
	}, nil
}
