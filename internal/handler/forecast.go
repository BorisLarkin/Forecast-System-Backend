package handler

import (
	"encoding/json"
	"net/http"
	"web/internal/ds"
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
func (h *Handler) JSONGetForecasts(ctx *gin.Context) {
	searchText := ctx.Query("search")
	var pred_len int
	var forec_empty bool
	user_id, _ := dsn.GetCurrentUserID()
	draft_id, err := h.Repository.GetUserDraftID(user_id)
	if err != nil {
		pred_len = 0
		draft_id = "none"
	} else {
		pred_len = h.Repository.GetPredLen(draft_id)
	}

	if searchText == "" {
		Forecasts, forec_len, err := h.Repository.ForecastList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		forec_empty = (forec_len == 0)
		ctx.JSON(http.StatusOK, gin.H{
			"Forecasts":   Forecasts,
			"forec_empty": forec_empty,
			"pred_len":    pred_len,
			"draft_id":    draft_id,
		})
	} else {
		filteredForecasts, forec_len, err := h.Repository.SearchForecast(searchText)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		forec_empty = (forec_len == 0)
		ctx.JSON(http.StatusOK, gin.H{
			"Forecasts":   filteredForecasts,
			"forec_empty": forec_empty,
			"pred_len":    pred_len,
			"draft_id":    draft_id,
		})
	}
}
func (h *Handler) JSONGetForecastById(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	var id string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	forecast, err := h.Repository.GetForecastByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}

func (h *Handler) JSONDeleteForecast(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	var id string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Minio.DeletePicture(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	err = h.Repository.DeleteForecast(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
func (h *Handler) JSONAddForecast(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	var forecast ds.Forecasts
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &forecast)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.CreateForecast(&forecast)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}

func (h *Handler) JSONEditForecast(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	var forecast ds.Forecasts
	var id string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &forecast)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Repository.EditForecast(&forecast, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}
func (h *Handler) JSONAddPicture(ctx *gin.Context) {
	jsonData, err := ctx.GetRawData()
	var id string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = json.Unmarshal(jsonData, &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	err = h.Minio.UploadPicture(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
