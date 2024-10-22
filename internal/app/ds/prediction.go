package ds

import "time"

type Predictions struct {
	Id                int       `json:"id" gorm:"primary_key"`
	Date_created      time.Time `json:"date_created" gorm:"not null"`
	Date_formed       time.Time `json:"date_formed"`
	Date_completed    time.Time `json:"date_completed"`
	UserID            int       `json:"-" gorm:"not null"`
	User              Users     `gorm:"foreignKey:UserID" json: "-" `
	ModerID           int       `json:"-"`
	Status            string    `gorm:"type:varchar(255); not null" json:"status"`
	Prediction_amount int       `json:"pred_amount"`
	Prediction_window int       `json:"pred_window"`
}

//status=deleted/draft/pending/done/error
