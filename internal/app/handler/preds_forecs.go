package handler

import (
	"web/internal/app/dsn"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddForecastToPred(ctx *gin.Context) {
	uid, err := dsn.GetCurrentUserID()
	if err != nil {
		return
	}
	f_id := ctx.Param("f_id")
	var pr_id string
	pr_id, err = h.Repository.GetUserDraftID(uid)
	if err != nil {
		h.Repository.CreateDraft()
		pr_id, _ = h.Repository.GetUserDraftID(uid)
	}
	h.Repository.CreatePreds_Forecs(pr_id, f_id)
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
	f_id := ctx.Param("f_id")
	h.Repository.DeletePreds_Forecs(pr_id, f_id)
}
