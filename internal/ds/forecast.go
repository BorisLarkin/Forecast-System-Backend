package ds

type Forecasts struct {
	Forecast_id   uint   `json:"forecast_id" gorm:"primary_key"`
	Title         string `gorm:"type:varchar(100); uniqueIndex; not null" json:"title"`
	Short         string `gorm:"type:varchar(50); uniqueIndex; not null" json:"short_title"`
	Descr         string `gorm:"type:varchar(255)" json:"descr"`
	Color         string `gorm:"type:varchar(50)" json:"color"`
	Img_url       string `gorm:"type:varchar(255)" json:"image"`
	Measure_type  string `gorm:"type:varchar(80); not null" json:"measure_type"`
	Extended_desc string `gorm:"type:varchar(1024)" json:"ext_desc"`
}

//http://127.0.0.1:9000/test/image-{{.Forecast_id}}.png

type ForecastRequest struct {
	Title         string `json:"title"`
	Short         string `json:"short_title"`
	Descr         string `json:"descr"`
	Color         string `json:"color"`
	Img_url       string `json:"image"`
	Measure_type  string `json:"measure_type"`
	Extended_desc string `json:"ext_desc"`
}

type GetForecastsResponse struct {
	Forecasts      *[]Forecasts `json:"forecasts"`
	DraftID        string       `json:"predicction_id"`
	DraftSize      int          `json:"prediction_size"`
	ForecastsEmpty bool         `json:"forecasts_empty"`
}

type ForecastResponse struct {
	ID            uint
	Title         string
	Short         string
	Descr         string
	Color         string
	Img_url       string
	Measure_type  string
	Extended_desc string
}

type ForecastResponseWithFlags struct {
	Forecast_id   uint   `json:"id"`
	Title         string `json:"title"`
	Short         string `json:"short_title"`
	Descr         string `json:"descr"`
	Color         string `json:"color"`
	Img_url       string `json:"image"`
	Measure_type  string `json:"measure_type"`
	Extended_desc string `json:"ext_desc"`
	Input         string `json:"input"`
	Result        string `json:"result"`
}
