package ds

type Preds_Forecs struct {
	PredictionID int    `json:"prediction_id" gorm:"primaryKey;auto_increment:false"`
	ForecastID   int    `json:"forecast_id" gorm:"primaryKey;auto_increment:false"`
	Input        string `gorm:"type:varchar(255)" json:"input"`
	Result       string `gorm:"type:varchar(255)" json:"result"`
}
