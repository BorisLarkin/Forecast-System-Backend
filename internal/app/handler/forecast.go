package handler

import (
	"net/http"
	"web/internal/app/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ForecastList(ctx *gin.Context) {
	forecastName := ctx.Query("search")
	var pred_len int
	user_id, _ := dsn.GetCurrentUserID()
	draft_id, _ := h.Repository.GetUserDraftID(user_id)

	if draft_id == "" {
		pred_len = 0
	} else {
		pred_len = h.Repository.GetPredLen(draft_id)
	}

	if forecastName == "" {
		Forecasts, err := h.Repository.ForecastList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts":    Forecasts,
			"jnz":          (pred_len == 0),
			"Curr_pred_id": draft_id,
			"Pred_len":     pred_len,
		})
	} else {
		filteredForecasts, err := h.Repository.SearchForecast(forecastName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.HTML(http.StatusOK, "forecasts.tmpl", gin.H{
			"Forecasts":    filteredForecasts,
			"jnz":          (pred_len == 0),
			"Curr_pred_id": draft_id,
			"Pred_len":     pred_len,
		})
	}
}
