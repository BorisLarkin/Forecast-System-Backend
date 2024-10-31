package handler

import (
	"web/internal/minio"
	"web/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Minio      *minio.MinioClient
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.MinioClient) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Minio:      m,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	admin := router.Group("/admin")
	router.GET("/forecasts", h.JSONGetForecasts)
	router.GET("/forecast/:id", h.JSONGetForecastById)
	admin.DELETE("/forecast/delete/:id", h.JSONDeleteForecast)
	admin.POST("/forecast/add", h.JSONAddForecast) //without img
	admin.PUT("/forecast/edit/:id", h.JSONEditForecast)
	router.POST("/add", h.JSONAddForecastToPred)
	admin.POST("/forecast/add_picture/:id", h.JSONAddPicture)

	router.GET("/predictions", h.JSONGetPredictions)                //status&form_data filter, no deleted or drafts
	router.GET("/prediction/:id", h.JSONGetPredictionById)          //+forecs
	router.DELETE("/prediction/delete/:id", h.JSONDeletePrediction) //form_data
	router.PUT("/prediction/form/:id", h.JSONFormPrediction)        //client-side
	router.PUT("/prediction/edit/:id", h.JSONEditPrediction)        //fields
	admin.PUT("/prediction/finish/:id", h.JSONFinishPrediction)     //decline or confirm + calc

	router.DELETE("/pr_fc/remove", h.JSONDeleteForecastFromPred)
	router.PUT("/pr_fc/edit", h.JSONEditPredForec)

	router.POST("/user/register", h.JSONUserRegister)
	router.PUT("/user/account/:id", h.JSONUserAccount)
	router.POST("/user/auth/:id", h.JSONUserAuth)
	router.POST("/user/deauth/:id", h.JSONUserDeauth)

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
