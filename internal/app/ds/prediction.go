package ds

import "time"

type Predictions struct {
	Id                int       `json:"id" gorm:"primary_key"`
	Date_created      time.Time `json:"date_created"`
	Date_formed       time.Time `json:"date_formed"`
	Date_completed    time.Time `json:"date_completed"`
	UserID            uint      `json:"-"`
	User              Users     `gorm:"foreignKey:UserID" json: "-" `
	ModerID           uint      `json:"-"`
	Status            string    `gorm:"type:varchar(255)" json:"status"`
	Prediction_amount int       `json:"pred_amount"`
	Prediction_window int       `json:"pred_window"`
}
