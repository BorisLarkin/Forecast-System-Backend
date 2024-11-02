package handler

import (
	"errors"
	"net/http"
	"time"
	"web/internal/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPredictions(ctx *gin.Context) {

	status := ctx.Query("status")
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	startDate, err := time.Parse("2006-Jan-02", startDateStr)
	if err != nil && startDateStr != "" {
		ctx.JSON(http.StatusBadRequest, "Invalid start date format")
		return
	}

	endDate, err := time.Parse("2006-Jan-02", endDateStr)
	if err != nil && endDateStr != "" {
		ctx.JSON(http.StatusBadRequest, "Invalid end date format")
		return
	}

	preds, err := h.Repository.GetPredictions(status, startDateStr != "", endDateStr != "", startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем результат клиенту
	ctx.JSON(http.StatusOK, preds)
}

func (h *Handler) GetPredictionById(ctx *gin.Context) {
	id := ctx.Param("id") // Получаем ID сообщения из URL

	prediction, err := h.Repository.GetPredictionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	forecs, err := h.Repository.GetForecastsByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"id":                prediction.Prediction_id,
		"status":            prediction.Status,
		"Prediction_amount": prediction.Prediction_amount,
		"Prediction_window": prediction.Prediction_window,
		"Date_created":      prediction.Date_created,
		"Date_formed":       prediction.Date_formed,
		"Date_completed":    prediction.Date_completed,
		"Creator":           prediction.UserID,
		"Moderator":         prediction.ModerID,
		"forecasts":         forecs,
	})
}

func (h *Handler) EditPrediction(ctx *gin.Context) {
	id := ctx.Param("id")

	var input struct {
		Amount int `json:"prediction_amount" binding:"required"`
		Window int `json:"prediction_window" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Repository.EditPrediction(id, input.Window, input.Amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем успешный ответ
	ctx.JSON(http.StatusOK, gin.H{"message": "Text updated successfully"})
}

func (h *Handler) FormPrediction(ctx *gin.Context) {
	id := ctx.Param("id")

	creatorID, err := dsn.GetCurrentUserID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Repository.FormPrediction(id, creatorID); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Status changed"})
}

func (h *Handler) FinishPrediction(ctx *gin.Context) {
	id := ctx.Param("id")

	is_admin, err := h.Repository.CurrentUser_IsAdmin()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	if is_admin != true {
		ctx.JSON(http.StatusConflict, errors.New("attempt to finish prediction as user"))
	}
	status := ctx.Query("status")
	if status == "complete" {
		//count
		if err := h.Repository.SetPredictionStatus(id, status); err != nil {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
	} else if status == "denied" {
		if err := h.Repository.SetPredictionStatus(id, status); err != nil {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
	} else {
		ctx.JSON(http.StatusConflict, errors.New("attempt to finish prediction with wrong status"))
	}
}

func (h *Handler) DeletePrediction(ctx *gin.Context) {
	pr_id := ctx.Param("id")

	creatorID, err := dsn.GetCurrentUserID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.Repository.DeletePrediction(pr_id, creatorID); err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Status changed successfully"})
}
