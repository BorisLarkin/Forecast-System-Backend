package handler

import (
	"fmt"
	"net/http"
	"web/internal/ds"
	"web/internal/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetForecasts(ctx *gin.Context) {
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
func (h *Handler) GetForecastById(ctx *gin.Context) {
	id := ctx.Param("id")

	forecast, err := h.Repository.GetForecastByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, forecast)
}

func (h *Handler) DeleteForecast(ctx *gin.Context) {
	id := ctx.Param("id")
	imageName := fmt.Sprintf("image-%s.png", id)
	err := h.Repository.DeletePicture(id, imageName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	err = h.Repository.DeleteForecast(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Forecast (id-%s) deleted", id)})
}
func (h *Handler) AddForecast(ctx *gin.Context) {
	var forecast ds.Forecasts

	if err := ctx.BindJSON(&forecast); err != nil {
		ctx.JSON(http.StatusBadRequest, "неверные данные")
		return
	}

	id, err := h.Repository.CreateForecast(&forecast)

	forecast.Forecast_id = id
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, forecast)
}

func (h *Handler) EditForecast(ctx *gin.Context) {
	var forecast ds.Forecasts
	id := ctx.Param("id")

	if err := ctx.BindJSON(&forecast); err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect JSON format")
		return
	}
	err := h.Repository.EditForecast(&forecast, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}
func (h *Handler) AddPicture(ctx *gin.Context) {
	forecast_id := ctx.Param("id")
	// Получаем файл изображения из запроса
	file, fileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to upload image")
		return
	}
	defer file.Close()

	imageName := fmt.Sprintf("image-%s.png", forecast_id)

	// Передаем файл в репозиторий для обработки
	err = h.Repository.UploadPicture(forecast_id, imageName, file, fileHeader.Size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image successfully uploaded"})
}
