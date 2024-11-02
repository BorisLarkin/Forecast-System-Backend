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
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.MinioClient) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	router.GET("/forecasts", h.GetForecasts)
	router.GET("/forecast/:id", h.GetForecastById)
	router.DELETE("/forecast/delete/:id", h.DeleteForecast)
	router.POST("/forecast/add", h.AddForecast) //without img
	router.PUT("/forecast/edit/:id", h.EditForecast)
	router.POST("/forecast_to_pred", h.AddForecastToPred)
	router.POST("/forecast/:id/add_picture", h.AddPicture)

	router.GET("/predictions", h.GetPredictions)                //status&form_data filter, no deleted or drafts
	router.GET("/prediction/:id", h.GetPredictionById)          //+forecs
	router.DELETE("/prediction/delete/:id", h.DeletePrediction) //form_data
	router.PUT("/prediction/form/:id", h.FormPrediction)        //client-side
	router.PUT("/prediction/edit/:id", h.EditPrediction)        //fields
	router.PUT("/prediction/finish/:id", h.FinishPrediction)    //decline or confirm + calc

	router.DELETE("/pr_fc/remove/:forecast_id/:prediction_id", h.DeleteForecastFromPred)
	router.PUT("/pr_fc/edit/:message_id/:chat_id", h.EditPredForec)

	router.POST("/user/register", h.UserRegister)
	router.PUT("/user/update/:id", h.UpdateUser)
	router.POST("/user/auth/:id", h.UserAuth)
	router.POST("/user/deauth/:id", h.UserDeauth)
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
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
