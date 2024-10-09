package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForecastList(ctx *gin.Context) {
	forecastName := ctx.Query("search")
	if forecastName == "" {
		Forecasts, err := h.Repository.ForecastList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts": Forecasts,
		})
	} else {
		filteredForecasts, err := h.Repository.SearchForecast(forecastName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecastss.tmpl", gin.H{
			"Forecasts": filteredForecasts,
		})
	}
}
