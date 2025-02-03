package ds

type Forecasts struct {
	Forecast_id   uint   `json:"forecast_id" gorm:"primary_key" binding:"required"`
	Title         string `gorm:"type:varchar(100); uniqueIndex; not null" json:"title" binding:"required"`
	Short         string `gorm:"type:varchar(50); uniqueIndex; not null" json:"short_title" binding:"required"`
	Descr         string `gorm:"type:varchar(255)" json:"descr" binding:"required"`
	Color         string `gorm:"type:varchar(50)" json:"color" binding:"required"`
	Img_url       string `gorm:"type:varchar(255)" json:"image"`
	Measure_type  string `gorm:"type:varchar(80); not null" json:"measure_type" binding:"required"`
	Extended_desc string `gorm:"type:varchar(1024)" json:"ext_desc" binding:"required"`
}

type ForecastRequest struct {
	Title         string `json:"title" binding:"required"`
	Short         string `json:"short_title" binding:"required"`
	Descr         string `json:"descr" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Img_url       string `json:"image"`
	Measure_type  string `json:"measure_type" binding:"required"`
	Extended_desc string `json:"ext_desc" binding:"required"`
}

type GetForecastsResponse struct {
	Forecasts      *[]Forecasts `json:"forecasts" binding:"required"`
	DraftID        string       `json:"prediction_id" binding:"required"`
	DraftSize      int          `json:"prediction_size" binding:"required"`
	ForecastsEmpty bool         `json:"forecasts_empty"`
}

type ForecastResponse struct {
	ID            uint   `json:"forecast_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Short         string `json:"short_title" binding:"required"`
	Descr         string `json:"descr" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Img_url       string `json:"image" binding:"required"`
	Measure_type  string `json:"measure_type" binding:"required"`
	Extended_desc string `json:"ext_desc" binding:"required"`
}

type ForecastResponseWithFlags struct {
	Forecast_id   uint   `json:"forecast_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Short         string `json:"short_title" binding:"required"`
	Descr         string `json:"descr" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Img_url       string `json:"image" binding:"required"`
	Measure_type  string `json:"measure_type" binding:"required"`
	Extended_desc string `json:"ext_desc" binding:"required"`
	Input         string `json:"input" binding:"required"`
	Result        string `json:"result" binding:"required"`
}
