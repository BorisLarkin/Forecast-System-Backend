package ds

type Preds_Forecs struct {
	Preds_forecs_id uint   `json:"preds_forecs_id" gorm:"primaryKey;auto_increment:false" binding:"required"`
	PredictionID    uint   `json:"prediction_id" gorm:"not null; uniqueIndex: pr_fc" binding:"required"`
	ForecastID      uint   `json:"forecast_id" gorm:"not null; uniqueIndex: pr_fc" binding:"required"`
	Input           string `gorm:"type:varchar(255)" json:"input"`
	Result          string `gorm:"type:varchar(255)" json:"result"`

	Prediction Predictions `gorm:"foreignKey:PredictionID"`
	Forecast   Forecasts   `gorm:"foreignKey:ForecastID"`
}

type UpdatePred_ForecInput struct {
	Input string `json:"input" binding:"required"`
}
