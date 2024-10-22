package ds

type Preds_Forecs struct {
	ID           int         `json:"id" gorm:"primaryKey;auto_increment:false"`
	PredictionID int         `json:"-" gorm:"not null; uniqueIndex: pr_fc"`
	Prediction   Predictions `json:"-" gorm:"foreignKey: PredictionID"`
	ForecastID   int         `json:"-" gorm:"not null; uniqueIndex: pr_fc"`
	Forecast     Forecasts   `json:"-" gorm:"foreignKey: ForecastID"`
	Input        string      `gorm:"type:varchar(255)" json:"input"`
	Result       string      `gorm:"type:varchar(255)" json:"result"`
}
