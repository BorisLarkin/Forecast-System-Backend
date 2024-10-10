package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PredictionById(ctx *gin.Context) {
	id := ctx.Param("id")
	Prediction, err := h.Repository.GetPredictionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	f, _, _ := h.Repository.GetForecastsByID(id)
	ctx.HTML(http.StatusOK, "prediction.tmpl", gin.H{
		"Prediction":   Prediction,
		"Pr_forecasts": f,
	})
}

func (h *Handler) DeletePrediction(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Repository.DeletePrediction(id)
	ctx.Redirect(http.StatusFound, "/forecasts")
}
