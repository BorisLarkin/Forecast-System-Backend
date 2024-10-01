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
	/*
		var Prediction_to_parse models.Prediction_parse
		Prediction_to_parse.Prediction = models.Predictions[slices.IndexFunc(models.Predictions, func(f models.Prediction) bool { return f.Id == pred_id })]
		Prediction_to_parse.Solid_cart_style = template.CSS("background-color: rgba" + color)
		Prediction_to_parse.Fade_cart_style = template.CSS("background-image:linear-gradient(90deg, rgba" + color + " 0%,rgba(255, 255, 255, 0) 100%)")

		for i, v := range models.Predictions {
			var Prediction_parse_tmp models.Prediction_parse
			Prediction_parse_tmp.Prediction = models.Predictions[i]
			Prediction_parse_tmp.Forecast = models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == v.F_id })]
			Prediction_parse_tmp.Solid_cart_style = template.CSS("background-color: rgba" + color)
			Prediction_parse_tmp.Fade_cart_style = template.CSS("background-image:linear-gradient(90deg, rgba" + color + " 0%,rgba(255, 255, 255, 0) 100%)")
			models.Prediction_parses = append(models.Prediction_parses, Prediction_parse_tmp)
		}
	*/
	c.HTML(http.StatusOK, "cart.tmpl", gin.H{
		"Predictions": models.Predictions,
		"Pred_header": models.HeaderDiv,
	})
}
