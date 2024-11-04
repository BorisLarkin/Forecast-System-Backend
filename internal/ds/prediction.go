package ds

import "time"

type Predictions struct {
	Prediction_id     int       `json:"prediction_id" gorm:"primary_key"`
	Date_created      time.Time `json:"date_created" gorm:"not null"`
	Date_formed       time.Time `json:"date_formed"`
	Date_completed    time.Time `json:"date_completed"`
	UserID            int       `json:"-" gorm:"not null"`
	User              Users     `json:"-" gorm:"foreignKey:UserID"`
	ModerID           int       `json:"-"`
	Status            string    `gorm:"type:varchar(255); check:status IN ('deleted', 'draft','pending','done','denied'); not null" json:"status"`
	Prediction_amount int       `json:"prediction_amount"`
	Prediction_window int       `json:"prediction_window"`
}

//status=deleted/draft/pending/complete/denied
