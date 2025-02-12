package ds

import (
	"time"
)

type Predictions struct {
	Prediction_id     uint      `json:"prediction_id" gorm:"primaryKey"`
	Date_created      time.Time `json:"date_created" gorm:"not null"`
	Date_formed       time.Time `json:"date_formed" gorm:"default:null"`
	Date_completed    time.Time `json:"date_completed" gorm:"default:null"`
	CreatorID         int       `json:"creator_id" gorm:"not null"`
	Creator           Users     `json:"-" gorm:"foreignKey:CreatorID"`
	ModerID           int       `json:"-"`
	Status            string    `gorm:"type:varchar(255); check:status IN ('deleted', 'draft','pending','done','denied'); not null" json:"status"`
	Prediction_amount int       `json:"prediction_amount"`
	Prediction_window int       `json:"prediction_window"`
	Qr                []byte    `json:"qr" gorm:"type:bytea;"`
}

type PredictionWithForecasts struct {
	Prediction Predictions                  `json:"prediction" binding:"required"`
	Forecasts  *[]ForecastResponseWithFlags `json:"forecasts" binding:"required"`
}

type PredictionWithUsers struct {
	Prediction Predictions `json:"prediction" binding:"required"`
	User       string      `json:"username" binding:"required"`
}
