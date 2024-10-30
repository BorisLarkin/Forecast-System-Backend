package handler

import (
	"net/http"
	"web/internal/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForecastList(ctx *gin.Context) {
	forecastName := ctx.Query("search")
	var pred_len int
	user_id, _ := dsn.GetCurrentUserID()
	draft_id, err := h.Repository.GetUserDraftID(user_id)

	if err != nil {
		pred_len = 0
		draft_id = "none"
	} else {
		pred_len = h.Repository.GetPredLen(draft_id)
	}

	if forecastName == "" {
		Forecasts, forec_len, err := h.Repository.ForecastList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts":    Forecasts,
			"forec_empty":  (forec_len == 0),
			"Curr_pred_id": draft_id,
			"Pred_len":     pred_len,
			"Search_str":   forecastName,
		})
	} else {
		filteredForecasts, forec_len, err := h.Repository.SearchForecast(forecastName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts":    filteredForecasts,
			"forec_empty":  (forec_len == 0),
			"Curr_pred_id": draft_id,
			"Pred_len":     pred_len,
			"Search_str":   forecastName,
		})
	}
}
