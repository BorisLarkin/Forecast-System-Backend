package handler

import (
	"fmt"
	"net/http"
	"web/internal/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddForecastToPred(ctx *gin.Context) {
	uid, err := dsn.GetCurrentUserID()
	if err != nil {
		return
	}
	f_id := ctx.Query("id")
	var pr_id string
	pr_id, err = h.Repository.GetUserDraftID(uid)
	if err != nil {
		h.Repository.CreateDraft()
		pr_id, _ = h.Repository.GetUserDraftID(uid)
	}
	h.Repository.CreatePreds_Forecs(pr_id, f_id)
	ctx.Redirect(http.StatusFound, "/forecasts")
}

func (h *Handler) DeleteForecastFromPred(ctx *gin.Context) {
	f_id := ctx.Param("forecast_id")
	pr_id := ctx.Param("prediction_id")
	h.Repository.DeletePreds_Forecs(pr_id, f_id)
	ln := h.Repository.GetPredLen(pr_id)
	user_id, err := dsn.GetCurrentUserID()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if ln == 0 {
		if err := h.Repository.DeletePrediction(pr_id, user_id); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		//ctx.Redirect(http.StatusFound, "/prediction/none")
	} else {
		//ctx.Redirect(http.StatusFound, "/prediction/"+pr_id)
	}
}

func (h *Handler) JSONAddForecastToPred(ctx *gin.Context) {
	uid, err := dsn.GetCurrentUserID()
	if err != nil {
		return
	}
	f_id := ctx.Query("id")
	var pr_id string
	pr_id, err = h.Repository.GetUserDraftID(uid)
	if err != nil {
		h.Repository.CreateDraft()
		pr_id, _ = h.Repository.GetUserDraftID(uid)
	}
	h.Repository.CreatePreds_Forecs(pr_id, f_id)
	ctx.Redirect(304, "/forecasts")
}

func (h *Handler) EditPredForec(ctx *gin.Context) {
	f_id := ctx.Param("forecast_id")
	pr_id := ctx.Param("prediction_id")

	sound, err := h.Repository.EditPredForec(uint(messageID), uint(chatID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Ошибка изменения поля 'со звуком'")
		return
	}

	// Возвращаем успешный ответ.
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Значение успешно изменено на '%v'", sound)})
}
