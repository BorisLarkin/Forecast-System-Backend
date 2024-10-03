package controllers

import (
	"net/http"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Parse_forecasts(c *gin.Context) {
	search := c.Query("search")
	var Fs []models.Forecast
	if search == "" {
		Fs = models.GetForecasts()
	} else {
		Fs = models.GetForecastsByName(search)
	}
	c.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
		"Forecasts":     Fs,
		"Forec_header":  models.HeaderDiv,
		"Curr_pred_len": models.GetCartLen(),
		"Curr_pred_id":  models.GetCurrPredictionId(),
	})
}
