package repository

import (
	"os"
	"web/internal/minio"

	"github.com/go-redis/redis/v8"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db          *gorm.DB
	logger      *logrus.Logger
	minioclient *minio.MinioClient
	redisClient *redis.Client
}

func New(dsn string, log *logrus.Logger, m *minio.MinioClient) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis options
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	return &Repository{
		db:          db,
		logger:      log,
		minioclient: m,
		redisClient: redisClient,
	}, nil
}
