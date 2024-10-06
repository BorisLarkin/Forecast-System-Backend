package ds

import "time"

type Prediction struct {
	Id             int `json:"id" gorm:"primary_key"`
	Date_created   time.Time
	Date_formed    time.Time
	Date_completed time.Time
	Creator        string
	Moderator      string
	Status         string
	Img_url        string `gorm:"type:varchar(255)" json:"img_url"`
	Title          string `gorm:"type:varchar(255)" json:"title"`
	Short          string `gorm:"type:varchar(255)" json:"short_title"`
	Desc           string `gorm:"type:varchar(255)" json:"desc"`
	Color          string `gorm:"type:varchar(255)" json:"color"`
	Measure_type   string `gorm:"type:varchar(255)" json:"measure_type"`
	Extended_desc  string `gorm:"type:varchar(255)" json:"ext_desc"`
}
