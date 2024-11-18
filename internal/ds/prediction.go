package ds

import (
	"time"
)

type Predictions struct {
	Prediction_id     uint      `json:"prediction_id" gorm:"primaryKey"`
	Date_created      time.Time `json:"date_created" gorm:"not null"`
	Date_formed       time.Time `json:"date_formed" gorm:"default:null"`
	Date_completed    time.Time `json:"date_completed" gorm:"default:null"`
	CreatorID         int       `json:"-" gorm:"not null"`
	Creator           Users     `json:"-" gorm:"foreignKey:CreatorID"`
	ModerID           int       `json:"-"`
	Status            string    `gorm:"type:varchar(255); check:status IN ('deleted', 'draft','pending','done','denied'); not null" json:"status"`
	Prediction_amount int       `json:"prediction_amount"`
	Prediction_window int       `json:"prediction_window"`
}

type PredictionDetail struct {
	ID                uint                         `json:"id"`
	Status            string                       `json:"status"`
	DateCreated       time.Time                    `json:"date_created"`
	DateFormed        time.Time                    `json:"date_formed,omitempty"`
	DateFinished      time.Time                    `json:"date_finished,omitempty"`
	Prediction_amount int                          `json:"prediction_amount"`
	Prediction_window int                          `json:"prediction_window"`
	Creator           Users                        `json:"creator"`
	Moderator         int                          `json:"moderator,omitempty"`
	Forecasts         *[]ForecastResponseWithFlags `json:"forecasts"`
}
