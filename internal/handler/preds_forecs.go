package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"web/internal/ds"

	"github.com/gin-gonic/gin"
)

// AddForecastToPred godoc
// @Summary      Add forecast to current user`s draft prediction
// @Description  If there`s no draft found, a new draft is to be created.
// @Tags         Preds_Forecs
// @Produce      json
// @Param        forecast_id path int true "Forecast ID"
// @Success      200
// @Failure      500
// @Router       /forecast_to_pred/{forecast_id} [post]
func (h *Handler) AddForecastToPred(ctx *gin.Context) {
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
	}
	uid := strconv.Itoa(int(payload.Uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "cannot get user id")
		return
	}
	f_id := ctx.Param("forecast_id")
	pr_id, err := h.Repository.GetUserDraftID(uid)
	if err != nil {
		h.Repository.CreateDraft(uid)
		pr_id, _ = h.Repository.GetUserDraftID(uid)
	}
	prediction, err := h.Repository.GetPredictionByID(pr_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "cannot get the prediction")
		return
	}
	if prediction.CreatorID != int(payload.Uid) || prediction.Status != "draft" {
		ctx.JSON(http.StatusBadRequest, "denied permission to edit the prediction")
		return
	}
	if err := h.Repository.CreatePreds_Forecs(pr_id, f_id); err != nil {
		ctx.JSON(http.StatusInternalServerError, "cannot create the record")
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("added forecast %s to prediction %s", f_id, pr_id))
}

/*
	func (h *Handler) AddForecastToDraft(ctx *gin.Context) {
		payload, err := h.GetTokenPayload(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
			return
		}
		uid := strconv.Itoa(int(payload.Uid))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "cannot get user id")
			return
		}
		f_id := ctx.Query("id")
		var pr_id string
		pr_id, err = h.Repository.GetUserDraftID(uid)
		if err != nil {
			h.Repository.CreateDraft(uid)
			pr_id, _ = h.Repository.GetUserDraftID(uid)
		}
		if err := h.Repository.CreatePreds_Forecs(pr_id, f_id); err != nil {
			ctx.JSON(http.StatusInternalServerError, "cannot create the record")
			return
		}
		pf, _ := h.Repository.GetPredForecByID(pr_id, f_id)
		ctx.JSON(http.StatusOK, pf)
		//ctx.Redirect(http.StatusFound, "/forecasts")
	}
*/

// DeleteForecastFromPred godoc
// @Summary      Delete forecast from specified prediction
// @Description  An error is returned in cases of unauthorized actions being attempted
// @Tags         Preds_Forecs
// @Produce      json
// @Param        forecast_id path int true "Forecast ID"
// @Param        prediction_id path int true "Prediction ID"
// @Success      200
// @Failure      400
// @Router       /pr_fc/remove/{prediction_id}/{forecast_id} [delete]
func (h *Handler) DeleteForecastFromPred(ctx *gin.Context) {
	f_id := ctx.Param("forecast_id")
	pr_id := ctx.Param("prediction_id")
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	uid := strconv.Itoa(int(payload.Uid))
	pred, err := h.Repository.GetPredictionByID(pr_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "prediction does not exist")
		return
	}
	if pred.CreatorID != int(payload.Uid) && payload.Role != ds.Moderator {
		ctx.JSON(http.StatusForbidden, "attempt to delete unowned prediction")
		return
	}
	if err := h.Repository.DeletePreds_Forecs(pr_id, f_id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ln := h.Repository.GetPredLen(pr_id)

	if ln == 0 {
		if err := h.Repository.DeletePrediction(pr_id, uid); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("pr_fc (%s, %s) deleted", f_id, pr_id)})
}

/*
	func (h *Handler) JSONAddForecastToPred(ctx *gin.Context) {
		payload, err := h.GetTokenPayload(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		}
		uid := strconv.Itoa(int(payload.Uid))

		f_id := ctx.Query("id")
		var pr_id string
		pr_id, err = h.Repository.GetUserDraftID(uid)
		if err != nil {
			h.Repository.CreateDraft(uid)
			pr_id, _ = h.Repository.GetUserDraftID(uid)
		}
		h.Repository.CreatePreds_Forecs(pr_id, f_id)
		ctx.Redirect(304, "/forecasts")
	}
*/

// EditPredForec godoc
// @Summary      Edit forecast`s input data for specified prediction
// @Description  An error is returned in cases of unauthorized actions being attempted
// @Tags         Preds_Forecs
// @Produce      json
// @Param        forecast_id path int true "Forecast ID"
// @Param        prediction_id path int true "Prediction ID"
// @Param        input body ds.UpdatePred_ForecInput true "New data"
// @Success      200
// @Failure      400
// @Router       /pr_fc/edit/{prediction_id}/{forecast_id} [delete]
func (h *Handler) EditPredForec(ctx *gin.Context) {
	f_id := ctx.Param("forecast_id")
	pr_id := ctx.Param("prediction_id")
	var input ds.UpdatePred_ForecInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	payload, err := h.GetTokenPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error retrieving token payload: %s", err))
		return
	}
	pred, err := h.Repository.GetPredictionByID(pr_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "prediction does not exist")
		return
	}
	if pred.CreatorID != int(payload.Uid) && payload.Role != ds.Moderator {
		ctx.JSON(http.StatusForbidden, "attempt to delete unowned prediction")
		return
	}

	if err := h.Repository.EditPredForec(f_id, pr_id, input.Input); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error changing input")
		return
	}

	// Возвращаем успешный ответ.
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successful change to '%s'", input.Input)})
}
