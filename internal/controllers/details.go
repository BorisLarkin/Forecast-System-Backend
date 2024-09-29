package controllers

import (
	"net/http"
	"slices"
	"strconv"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Details_parse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	Forecast_to_det := models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == id })]
	c.HTML(http.StatusOK, "details.tmpl", gin.H{
		"Forecast_to_det": Forecast_to_det,
		"Det_header":      models.HeaderDiv,
	})
}
