package repository

import (
	"web/internal/minio"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	logger      *logrus.Logger
	minioclient *minio.MinioClient
}

func New(dsn string, log *logrus.Logger, m *minio.MinioClient) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:          db,
		logger:      log,
		minioclient: m,
	}, nil
}
