package controllers

import (
	//"html/template"
	"net/http"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Parse_menu(c *gin.Context) {
	search := c.Query("search")
	var Fs []models.Forecast
	if search == "" {
		values := c.Request.URL.Query()
		values.Del("search")
		c.Request.URL.RawQuery = values.Encode()
		c.Next()
		Fs = models.GetForecasts()
	} else {
		Fs = models.GetForecastsByName(search)
	}
	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"Forecasts":     Fs,
		"Forec_header":  models.HeaderDiv,
		"Curr_pred_len": models.GetCartLen(),
		"Curr_pred_id":  models.GetCurrPredictionId(),
	})
}
