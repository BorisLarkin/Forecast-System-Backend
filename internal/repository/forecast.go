package repository

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"web/internal/ds"

	"github.com/minio/minio-go/v7"
)

func (r *Repository) ForecastList() (*[]ds.Forecasts, int, error) {
	var Forecasts []ds.Forecasts
	r.db.Find(&Forecasts)
	return &Forecasts, len(Forecasts), nil
}

func (r *Repository) GetForecastByID(id string) (*ds.Forecasts, error) {
	var Forecast ds.Forecasts
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Forecast, intId)
	return &Forecast, nil
}
func (r *Repository) SearchForecast(search string) (*[]ds.Forecasts, int, error) {
	var Forecast []ds.Forecasts
	r.db.Find(&Forecast)

	var filteredForecast []ds.Forecasts
	for _, f := range Forecast {
		if strings.Contains(strings.ToLower(f.Title), strings.ToLower(search)) {
			filteredForecast = append(filteredForecast, f)
		}
	}
	return &filteredForecast, len(filteredForecast), nil
}

func (r *Repository) CreateForecast(forecast *ds.Forecasts) (uint, error) {
	err := r.db.Create(&forecast).Error
	if err != nil {
		return 0, fmt.Errorf("error creating forecast: %w", err)
	}
	return forecast.Forecast_id, nil
}

func (r *Repository) DeleteForecast(id string) error {
	err := r.db.Delete(&ds.Forecasts{}, id).Error
	if err != nil {
		return fmt.Errorf("error deleting forecast with id %s: %w", id, err)
	}
	return nil
}
func (r *Repository) EditForecast(forecast *ds.Forecasts) error {
	err := r.db.Save(&forecast).Error
	if err != nil {
		return fmt.Errorf("error editing forecast (%d): %w", forecast.Forecast_id, err)
	}
	return nil
}

func (r *Repository) DeletePicture(id string, img string) error {
	err := r.minioclient.RemoveObject(context.Background(), "test", img, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("img deletion error")
	}
	return nil
}

func (r *Repository) UploadPicture(id string, imageName string, imageFile io.Reader, imageSize int64) error {
	var forecast ds.Forecasts

	// Найти чат по ID
	if err := r.db.First(&forecast, id).Error; err != nil {
		return fmt.Errorf("forecast (%s) not found: %w", id, err)
	}

	// Если старое изображение существует, удалить его из Minio
	if forecast.Img_url != "" {
		err := r.minioclient.RemoveObject(context.Background(), "test", imageName, minio.RemoveObjectOptions{})
		if err != nil {
			return fmt.Errorf("img delete error %s: %v", forecast.Img_url, err)
		}
	}

	// Загрузить новое изображение в Minio
	_, errMinio := r.minioclient.PutObject(context.Background(), "test", imageName, imageFile, imageSize, minio.PutObjectOptions{
		ContentType: "image/png",
	})

	forecast.Img_url = fmt.Sprintf("http://127.0.0.1:9000/test/image-%s.png", id)
	errDB := r.db.Save(&forecast).Error

	if errMinio != nil || errDB != nil {
		return fmt.Errorf("img upload error for forecast %s", id)
	}

	return nil
}
