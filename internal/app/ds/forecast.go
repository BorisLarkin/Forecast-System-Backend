package ds

type Forecasts struct {
	Id            int    `json:"forecast_id" gorm:"primary_key"`
	Title         string `gorm:"type:varchar(100); uniqueIndex; not null" json:"title"`
	Short         string `gorm:"type:varchar(50); uniqueIndex; not null" json:"short_title"`
	Descr         string `gorm:"type:varchar(255)" json:"descr"`
	Color         string `gorm:"type:varchar(50)" json:"color"`
	Measure_type  string `gorm:"type:varchar(80); not null" json:"measure_type"`
	Extended_desc string `gorm:"type:varchar(1024)" json:"ext_desc"`
}
