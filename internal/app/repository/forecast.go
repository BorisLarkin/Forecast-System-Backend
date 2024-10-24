package repository

import (
	"strconv"
	"strings"
	"web/internal/app/ds"
)

func (r *Repository) ForecastList() (*[]ds.Forecasts, int, error) {
	var Forecasts []ds.Forecasts
	r.db.Find(&Forecasts)
	return &Forecasts, len(Forecasts), nil
}

func (r *Repository) GetForecastByID(id string) (*ds.Forecasts, error) {
	var Forecast ds.Forecasts
	intId, _ := strconv.Atoi(id)
	r.db.Find(&Forecast, intId)
	return &Forecast, nil
}
func (r *Repository) SearchForecast(search string) (*[]ds.Forecasts, int, error) {
	var Forecast []ds.Forecasts
	r.db.Find(&Forecast)

	var filteredForecast []ds.Forecasts
	for _, f := range Forecast {
		if strings.Contains(strings.ToLower(f.Short), strings.ToLower(search)) {
			filteredForecast = append(filteredForecast, f)
		}
	}
	return &filteredForecast, len(filteredForecast), nil
}

func (r *Repository) CreateForecast(forecast *ds.Forecasts) error {
	return r.db.Create(forecast).Error
}

func (r *Repository) DeleteForecast(id string) {
	query := "DELETE FROM forecasts WHERE forecast_id = ?"
	r.db.Exec(query, id)
}
