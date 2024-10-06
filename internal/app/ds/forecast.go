package ds

type Forecast struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	Img_url       string `gorm:"type:varchar(255)" json:"img_url"`
	Title         string `gorm:"type:varchar(255)" json:"title"`
	Short         string `gorm:"type:varchar(255)" json:"short_title"`
	Desc          string `gorm:"type:varchar(255)" json:"desc"`
	Color         string `gorm:"type:varchar(255)" json:"color"`
	Measure_type  string `gorm:"type:varchar(255)" json:"measure_type"`
	Extended_desc string `gorm:"type:varchar(255)" json:"ext_desc"`
}
