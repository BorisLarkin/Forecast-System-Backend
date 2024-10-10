package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForecastList(ctx *gin.Context) {
	forecastName := ctx.Query("search")
	jnz := false
	if forecastName == "" {
		Forecasts, err := h.Repository.ForecastList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(*Forecasts) == 0 {
			jnz = true
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts":    Forecasts,
			"len_jnz":      jnz,
			"Curr_pred_id": 1,
		})
	} else {
		filteredForecasts, err := h.Repository.SearchForecast(forecastName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(*filteredForecasts) == 0 {
			jnz = true
		}
		ctx.HTML(http.StatusOK, "forecastss.tmpl", gin.H{
			"Forecasts":    filteredForecasts,
			"len_jnz":      jnz,
			"Curr_pred_id": 1,
		})
	}
}
