package controllers

import (
	"net/http"
	"strconv"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func ParseCart(c *gin.Context) {
	id := c.Param("id")
	pred_id, err := strconv.Atoi(id)
	if (err) != nil || pred_id < 0 || pred_id >= len(models.Predictions) {
		c.String(http.StatusNotFound, "Страница не найдена")
		return
	}
	var Prediction_to_parse models.Prediction = models.GetPredictionById(pred_id)
	var Forecasts_to_parse []models.Forecast = models.GetForecastsByPredictionId(pred_id)

	c.HTML(http.StatusOK, "cart.tmpl", gin.H{
		"Prediction_to_parse": Prediction_to_parse,
		"Forecasts_to_parse":  Forecasts_to_parse,
		"Pred_header":         models.HeaderDiv,
	})
}
