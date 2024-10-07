package ds

import "time"

type Predictions struct {
	Id             int       `json:"id" gorm:"primary_key"`
	Date_created   time.Time `json:"date_created"`
	Date_formed    time.Time `json:"date_formed"`
	Date_completed time.Time `json:"date_completed"`
	User           Users     `gorm:"foreignKey:UserID" json: "-"`
	UserID         uint      `json:"-"`
	ModerID        uint      `json:"-"`
	Status         string    `gorm:"type:varchar(255)" json:"status"`
	Input          string    `gorm:"type:varchar(255)" json:"input"`
	Result         string    `gorm:"type:varchar(255)" json:"result"`
}
