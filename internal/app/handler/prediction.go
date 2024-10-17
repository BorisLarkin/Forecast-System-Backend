package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PredictionById(ctx *gin.Context) {
	id := ctx.Param("id")
	is_init := true
	if id == "none" {
		is_init = false
		ctx.HTML(http.StatusOK, "prediction.tmpl", gin.H{
			"Exists": is_init,
		})
	} else {
		Prediction, err := h.Repository.GetPredictionByID(id)
		if err != nil {
			is_init = false
			ctx.HTML(http.StatusOK, "prediction.tmpl", gin.H{
				"Exists": is_init,
			})
		}
		is_draft := false
		if Prediction.Status == "draft" {
			is_draft = true
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		f, _ := h.Repository.GetForecastsByID(id)
		ctx.HTML(http.StatusOK, "prediction.tmpl", gin.H{
			"Prediction":   Prediction,
			"Pr_forecasts": f,
			"Exists":       is_init,
			"Draft":        is_draft,
		})
	}
}

func (h *Handler) DeletePrediction(ctx *gin.Context) {
	id := ctx.Query("id")
	h.Repository.DeletePrediction(id)
	ctx.Redirect(http.StatusFound, "/forecasts")
}

func (h *Handler) SavePrediction(ctx *gin.Context) {
	id := ctx.Query("id")
	h.Repository.DeletePrediction(id)
	ctx.Redirect(http.StatusFound, "/forecasts")
}
