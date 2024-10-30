package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DetailsById(ctx *gin.Context) {
	id := ctx.Query("id")
	Forecast, err := h.Repository.GetForecastByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "details.tmpl", gin.H{
		"Forecast_to_det": Forecast,
	})
}
