package controllers

import (
	"html/template"
	"net/http"
	"slices"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func ParseCart(c *gin.Context) {
	models.Prediction_parses = models.Prediction_parses[:0]
	for i, v := range models.Predictions {
		var Prediction_parse_tmp models.Prediction_parse
		Prediction_parse_tmp.Prediction = models.Predictions[i]
		Prediction_parse_tmp.Forecast = models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == v.F_id })]
		var color = Prediction_parse_tmp.Forecast.Color
		Prediction_parse_tmp.Solid_cart_style = template.CSS("background-color: rgba" + color)
		Prediction_parse_tmp.Fade_cart_style = template.CSS("background-image:linear-gradient(90deg, rgba" + color + " 0%,rgba(255, 255, 255, 0) 100%)")
		models.Prediction_parses = append(models.Prediction_parses, Prediction_parse_tmp)
	}
	c.HTML(http.StatusOK, "cart.tmpl", gin.H{
		"Predictions": models.Prediction_parses,
	})
}
