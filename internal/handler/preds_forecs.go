package handler

import (
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
	uid, err := dsn.GetCurrentUserID()
	if err != nil {
		return
	}
	pr_id, err := h.Repository.GetUserDraftID(uid)
	if err != nil {
		return
	}
	f_id := ctx.Query("id")
	h.Repository.DeletePreds_Forecs(pr_id, f_id)
	ln := h.Repository.GetPredLen(pr_id)
	if ln == 0 {
		h.Repository.DeletePrediction(pr_id)
		ctx.Redirect(http.StatusFound, "/prediction/none")
	} else {
		ctx.Redirect(http.StatusFound, "/prediction/"+pr_id)
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
