package handler

import (
	"web/internal/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
}

func NewHandler(l *logrus.Logger, r *repository.Repository) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/forecasts", h.ForecastList)
	router.GET("/prediction/:id", h.PredictionById)
	router.GET("/details", h.DetailsById)
	router.POST("/statuschange", h.HandleStatusChange)
	router.POST("/add", h.AddForecastToPred)
	//router.POST("/remove", h.DeleteForecastFromPred)
	registerStatic(router)
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/templates/*")
	router.Static("/static", "./static")
}

func (h *Handler) HandleStatusChange(ctx *gin.Context) {
	operation := ctx.Query("operation")
	if operation == "delete" {
		h.DeleteDraft(ctx)
		return
	}
	if operation == "save" {
		h.SavePrediction(ctx)
	}
}
