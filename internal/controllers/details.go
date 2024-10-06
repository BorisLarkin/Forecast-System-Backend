package controllers

import (
	"net/http"
	"slices"
	"strconv"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Details_parse(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		intId, _ := strconv.Atoi(id)
		Forecast_to_det := models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == intId })]
		c.HTML(http.StatusOK, "details.tmpl", gin.H{
			"Forecast_to_det": Forecast_to_det,
			"Det_header":      models.HeaderDiv,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "try with id",
	})
}
