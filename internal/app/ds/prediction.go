package ds

import "time"

type Predictions struct {
	Id                int       `json:"id" gorm:"primary_key"`
	Date_created      time.Time `json:"date_created"`
	Date_formed       time.Time `json:"date_formed"`
	Date_completed    time.Time `json:"date_completed"`
	UserID            int       `json:"-"`
	User              Users     `gorm:"foreignKey:UserID" json: "-" `
	ModerID           int       `json:"-"`
	Status            string    `gorm:"type:varchar(255)" json:"status"`
	Prediction_amount int       `json:"pred_amount"`
	Prediction_window int       `json:"pred_window"`
}

//status=deleted/draft/in-work/done/recieved
