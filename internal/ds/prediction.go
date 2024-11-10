package ds

import "time"

type Predictions struct {
	Prediction_id     uint      `json:"prediction_id" gorm:"primaryKey"`
	Date_created      time.Time `json:"date_created" gorm:"not null"`
	Date_formed       time.Time `json:"date_formed" gorm:"default:null"`
	Date_completed    time.Time `json:"date_completed" gorm:"default:null"`
	CreatorID         int       `json:"-" gorm:"not null"`
	Creator           Users     `json:"-" gorm:"foreignKey:CreatorID"`
	ModerID           int       `json:"-"`
	Moder             Users     `json:"-" gorm:"foreignKey:ModerID"`
	Status            string    `gorm:"type:varchar(255); check:status IN ('deleted', 'draft','pending','done','denied'); not null" json:"status"`
	Prediction_amount int       `json:"prediction_amount"`
	Prediction_window int       `json:"prediction_window"`
}

//status=deleted/draft/pending/complete/denied

type PredictionWithUsers struct {
	Prediction_id     uint      `json:"prediction_id"`
	Date_created      time.Time `json:"date_created"`
	Date_formed       time.Time `json:"date_formed,omitempty"`
	Date_completed    time.Time `json:"date_completed,omitempty"`
	Status            string    `json:"status"`
	Prediction_amount int       `json:"prediction_amount"`
	Prediction_window int       `json:"prediction_window"`
	Creator           string    `json:"creator"`
	Moderator         string    `json:"moderator,omitempty"`
}

type PredictionDetail struct {
	ID                uint                        `json:"id"`
	Status            string                      `json:"status"`
	DateCreate        time.Time                   `json:"date_create"`
	DateUpdate        time.Time                   `json:"date_update,omitempty"`
	DateFinish        time.Time                   `json:"date_finish,omitempty"`
	Prediction_amount int                         `json:"prediction_amount"`
	Prediction_window int                         `json:"prediction_window"`
	Creator           string                      `json:"creator"`
	Moderator         string                      `json:"moderator,omitempty"`
	Forecasts         []ForecastResponseWithFlags `json:"chats"`
}
