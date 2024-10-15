package ds

type Preds_Forecs struct {
	ID           int    `json:"id" gorm:"primaryKey;auto_increment:false"`
	PredictionID int    `json:"prediction_id" gorm:"uniqueIndex:pr_fc"`
	ForecastID   int    `json:"forecast_id" gorm:"uniqueIndex:pr_fc"`
	Input        string `gorm:"type:varchar(255)" json:"input"`
	Result       string `gorm:"type:varchar(255)" json:"result"`
}
