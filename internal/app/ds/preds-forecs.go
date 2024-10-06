package ds

type Preds_Forecs struct {
	PredictionID uint `json:"prediction_id" gorm:"primaryKey;auto_increment:false"`
	ForecastID   uint `json:"forecast_id" gorm:"primaryKey;auto_increment:false"`
}
