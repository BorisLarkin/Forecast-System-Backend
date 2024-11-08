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
	router.POST("/forecast_to_pred/:forecast_id", h.AddForecastToPred)
	router.POST("/forecast/:id/add_picture", h.AddPicture)

	router.GET("/predictions", h.GetPredictions)                //status&form_data filter, no deleted or drafts
	router.GET("/prediction/:id", h.GetPredictionById)          //+forecs
	router.DELETE("/prediction/delete/:id", h.DeletePrediction) //form_data
	router.PUT("/prediction/form/:id", h.FormPrediction)        //client-side
	router.PUT("/prediction/edit/:id", h.EditPrediction)        //fields
	router.PUT("/prediction/finish/:id", h.FinishPrediction)    //decline or confirm + calc

	router.DELETE("/pr_fc/remove/:prediction_id/:forecast_id", h.DeleteForecastFromPred)
	router.PUT("/pr_fc/edit/:prediction_id/:forecast_id", h.EditPredForec)

	router.POST("/user/register", h.RegisterUser)
	router.PUT("/user/update/:id", h.UpdateUser)
	router.POST("/user/login", h.LoginUser)
	router.POST("/user/logout", h.LogoutUser)
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/templates/*")
	router.Static("/static", "./static")
}
