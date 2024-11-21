package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
)

// GetPredictions godoc
// @Summary      Show all predictions made for current user
// @Description  very very friendly response
// @Tags         Predictions
// @Produce      json
// @Param status query string false "Prediction status"
// @Param start_date query string false "Earliest date created filter: YYYY-Mon-DD"
// @Param end_date query string false "Latest date created filter: YYYY-Mon-DD"
// @Success      200  {object}  []ds.Predictions
// @Failure      400
// @Router       /predictions [get]
func (h *Handler) GetPredictions(ctx *gin.Context) {
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
	}

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
	uid_string := strconv.Itoa(int(payload.Uid))
	preds, err := h.Repository.GetPredictions(uid_string, payload.Role, status, startDateStr != "", endDateStr != "", startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, preds)
}

// GetPredictionByID godoc
// @Summary      Display a prediction and its forecasts
// @Description  very very friendly response
// @Tags         Predictions
// @Produce      json
// @Param        id path int true "Prediction ID"
// @Success      200  {object}  ds.PredictionWithForecasts
// @Failure      403
// @Router       /prediction/{id} [get]
func (h *Handler) GetPredictionById(ctx *gin.Context) {
	id := ctx.Param("id")

	prediction, err := h.Repository.GetPredictionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	if prediction.CreatorID != int(payload.Uid) && payload.Role != ds.Moderator {
		ctx.JSON(http.StatusForbidden, fmt.Errorf("attempt to view unowned prediction"))
		return
	}
	forecs, err := h.Repository.GetForecastsByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, ds.PredictionWithForecasts{
		Prediction: *prediction,
		Forecasts:  forecs,
	})
}

type EditPredReq struct {
	Amount int `json:"prediction_amount" binding:"required"`
	Window int `json:"prediction_window" binding:"required"`
}

// EditPrediction godoc
// @Summary      Edit specified prediction`s prediction amount & window
// @Description  very very friendly response
// @Tags         Predictions
// @Accept       json
// @Produce      json
// @Param        prediction body EditPredReq true "New prediction data"
// @Param        id path int true "Prediction ID"
// @Success      200 {object}  ds.PredictionWithForecasts
// @Failure      403
// @Router       /prediction/edit/{id} [put]
func (h *Handler) EditPrediction(ctx *gin.Context) {
	id := ctx.Param("id")
	var input EditPredReq
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	prediction, err := h.Repository.GetPredictionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	if prediction.CreatorID != int(payload.Uid) && payload.Role != ds.Moderator {
		ctx.JSON(http.StatusForbidden, fmt.Errorf("attempt to view unowned prediction"))
	}
	pred, err := h.Repository.EditPrediction(id, input.Window, input.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	forecs, _ := h.Repository.GetForecastsByID(string(pred.Prediction_id))
	ctx.JSON(http.StatusOK, ds.PredictionWithForecasts{
		Prediction: *pred,
		Forecasts:  forecs,
	})
}

// FormPrediction godoc
// @Summary      Form specified prediction
// @Description  very very friendly response
// @Tags         Predictions
// @Produce      json
// @Param        id path int true "Prediction ID"
// @Success      200 {object}  ds.PredictionWithForecasts
// @Failure      403
// @Router       /prediction/form/{id} [put]
func (h *Handler) FormPrediction(ctx *gin.Context) {
	id := ctx.Param("id")
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	creatorID := strconv.Itoa(int(payload.Uid))

	if err := h.Repository.FormPrediction(id, creatorID); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	pred, _ := h.Repository.GetPredictionByID(id)
	forecs, _ := h.Repository.GetForecastsByID(id)
	ctx.JSON(http.StatusOK, ds.PredictionWithForecasts{
		Prediction: *pred,
		Forecasts:  forecs,
	})
}

// FinishPrediction godoc
// @Summary      Finish specified prediction
// @Description  Can be ended with statuses: ["denied", "completed"]
// @Tags         Predictions
// @Produce      json
// @Param        id path int true "Prediction ID"
// @Param status query string false "Status to be set"
// @Success      200 {object}  ds.PredictionWithForecasts
// @Failure      409
// @Router       /prediction/finish/{id} [put]
func (h *Handler) FinishPrediction(ctx *gin.Context) {
	id := ctx.Param("id")
	status := ctx.Query("status")

	if status == "completed" {
		_, err := h.Repository.CalculatePrediction(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := h.Repository.SetPredictionStatus(id, "done"); err != nil {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
		pred, _ := h.Repository.GetPredictionByID(id)
		forecs, _ := h.Repository.GetForecastsByID(id)
		ctx.JSON(http.StatusOK, ds.PredictionWithForecasts{
			Prediction: *pred,
			Forecasts:  forecs,
		})
	} else if status == "denied" {
		if err := h.Repository.SetPredictionStatus(id, status); err != nil {
			ctx.JSON(http.StatusConflict, err.Error())
			return
		}
		pred, _ := h.Repository.GetPredictionByID(id)
		forecs, _ := h.Repository.GetForecastsByID(id)
		ctx.JSON(http.StatusOK, ds.PredictionWithForecasts{
			Prediction: *pred,
			Forecasts:  forecs,
		})
	} else {
		ctx.JSON(http.StatusConflict, errors.New("attempt to finish prediction with wrong status"))
	}
}

// DeletePrediction godoc
// @Summary      Delete specified prediction
// @Description  Method sets prediction`s status to "deleted" without actually removing it from the db model
// @Tags         Predictions
// @Produce      json
// @Param        id path int true "Prediction ID"
// @Success      200
// @Failure      403
// @Router       /prediction/delete/{id} [delete]
func (h *Handler) DeletePrediction(ctx *gin.Context) {
	pr_id := ctx.Param("id")
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
	}
	creatorID := strconv.Itoa(int(payload.Uid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.Repository.DeletePrediction(pr_id, creatorID); err != nil {
		ctx.JSON(http.StatusForbidden, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Status changed successfully"})
}

// nil
func (h *Handler) CreateDraft(ctx *gin.Context) {
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
	}
	creatorID := strconv.Itoa(int(payload.Uid))
	if err := h.Repository.CreateDraft(creatorID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not create draft"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "draft created successfully"})
}
