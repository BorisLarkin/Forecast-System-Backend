package ds

type Preds_Forecs struct {
	Preds_forecs_id int    `json:"preds_forecs_id" gorm:"primaryKey;auto_increment:false"`
	PredictionID    int    `json:"prediction_id" gorm:"not null; uniqueIndex: pr_fc"`
	ForecastID      int    `json:"forecast_id" gorm:"not null; uniqueIndex: pr_fc"`
	Input           string `gorm:"type:varchar(255)" json:"input"`
	Result          string `gorm:"type:varchar(255)" json:"result"`
}

type Forecs_inputs struct {
	Forecast Forecasts
	Input    string
}
