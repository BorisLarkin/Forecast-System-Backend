package ds

type Forecasts struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Title         string `gorm:"type:varchar(255)" json:"title"`
	Short         string `gorm:"type:varchar(255)" json:"short_title"`
	Descr         string `gorm:"type:varchar(255)" json:"descr"`
	Color         string `gorm:"type:varchar(255)" json:"color"`
	Measure_type  string `gorm:"type:varchar(255)" json:"measure_type"`
	Extended_desc string `gorm:"type:varchar(1023)" json:"ext_desc"`
}
