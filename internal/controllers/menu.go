package controllers

import (
	"html/template"
	"net/http"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Parse_menu(c *gin.Context) {
	models.Forecast_parses = models.Forecast_parses[:0]
	for i, v := range models.Forecasts {
		var Forecast_parse_tmp models.Forecast_parse
		Forecast_parse_tmp.Forecast = v
		Forecast_parse_tmp.Solid_style = template.CSS("background-color: rgba" + models.Forecasts[i].Color)
		Forecast_parse_tmp.Fade_style = template.CSS("background-image:linear-gradient(0deg, rgba" + models.Forecasts[i].Color + " 0%,rgba(255, 255, 255, 0) 100%)")
		models.Forecast_parses = append(models.Forecast_parses, Forecast_parse_tmp)
	}

	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"Forecasts": models.Forecast_parses,
	})
}
