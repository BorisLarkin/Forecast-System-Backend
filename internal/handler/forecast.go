package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
)

// GetForecasts godoc
// @Summary      Show all available forecasts filtered by name
// @Description  very very friendly response
// @Tags         Forecasts
// @Produce      json
// @Param searchText query string false "name filter"
// @Success      200  {object}  ds.GetForecastsResponse
// @Failure      500
// @Router       /forecasts [get]
func (h *Handler) GetForecasts(ctx *gin.Context) {
	searchText := ctx.Query("forecast_name")
	var pred_len int
	var forec_empty bool
	var draft_id string
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		pred_len = 0
		draft_id = "none"
	} else {
		uid := strconv.Itoa(int(payload.Uid))
		draft_id, err = h.Repository.GetUserDraftID(uid)
		if err != nil {
			pred_len = 0
			draft_id = "none"
		} else {
			pred_len = h.Repository.GetPredLen(draft_id)
		}
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
		ctx.JSON(http.StatusOK, ds.GetForecastsResponse{
			Forecasts:      Forecasts,
			DraftID:        draft_id,
			DraftSize:      pred_len,
			ForecastsEmpty: forec_empty,
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
		ctx.JSON(http.StatusOK, ds.GetForecastsResponse{
			Forecasts:      filteredForecasts,
			DraftID:        draft_id,
			DraftSize:      pred_len,
			ForecastsEmpty: forec_empty,
		})
	}
}

// GetForecastByID godoc
// @Summary      Get a specified forecast by its ID
// @Description  very very friendly response
// @Tags         Forecasts
// @Produce      json
// @Param        id path int true "Forecast ID"
// @Success      200  {object}  ds.ForecastResponse
// @Failure      404
// @Router       /forecast/{id} [get]
func (h *Handler) GetForecastById(ctx *gin.Context) {
	id := ctx.Param("id")

	forecast, err := h.Repository.GetForecastByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, forecast)
}

// DeleteForecast godoc
// @Summary      Delete a specified forecast by its ID
// @Description  very very friendly response
// @Tags         Forecasts
// @Produce      json
// @Param        id path int true "Forecast ID"
// @Success      200
// @Failure      400
// @Router       /forecast/delete/{id} [delete]
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
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Forecast (id-%s) deleted", id)})
}

// AddForecast godoc
// @Summary      Add forecast to the list
// @Description  very very friendly response
// @Tags         Forecasts
// @Accept       json
// @Produce      json
// @Param        forecast body ds.ForecastRequest true "New forecast data"
// @Success      200  {object}  ds.ForecastRequest
// @Failure      500  {object}  ds.Forecasts
// @Router       /forecast/add [post]
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

// EditForecast godoc
// @Summary      Edit forecast
// @Description  very very friendly response
// @Tags         Forecasts
// @Accept       json
// @Produce      json
// @Param id path int true "Forecast ID"
// @Param        forecast body ds.ForecastRequest true "New forecast data"
// @Success      200  {object}  ds.ForecastRequest
// @Failure      500  {object}  ds.Forecasts
// @Router       /forecast/edit/{id} [put]
func (h *Handler) EditForecast(ctx *gin.Context) {
	var forecast ds.Forecasts
	id := ctx.Param("id")

	if err := ctx.BindJSON(&forecast); err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect JSON format")
		return
	}
	intid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect id format")
		return
	}
	forecast.Forecast_id = uint(intid)
	err = h.Repository.EditForecast(&forecast)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}

// AddImageToForecast godoc
// @Summary      Add image to specified forecast
// @Description  very very friendly response
// @Tags         Forecasts
// @Accept multipart/form-data
// @Param image formData file true "New image for the forecast"
// @Param id path int true "Forecast ID"
// @Success      200
// @Failure      500
// @Router       /forecast/{id}/add_picture [post]
func (h *Handler) AddPicture(ctx *gin.Context) {
	forecast_id := ctx.Param("id")
	// Get file out of the body
	file, fileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to upload image", "error": err.Error()})
		return
	}
	defer file.Close()

	imageName := fmt.Sprintf("%s.png", forecast_id)

	err = h.Repository.UploadPicture(forecast_id, imageName, file, fileHeader.Size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image successfully uploaded"})
}
