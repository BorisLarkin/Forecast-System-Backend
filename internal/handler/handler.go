package handler

import (
	"web/internal/config"
	"web/internal/minio"
	"web/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	Logger     *logrus.Logger
	Repository *repository.Repository
	Config     *config.Config
}

func NewHandler(l *logrus.Logger, r *repository.Repository, m *minio.MinioClient, c *config.Config) *Handler {
	return &Handler{
		Logger:     l,
		Repository: r,
		Config:     c,
	}
}

func (h *Handler) RegisterHandler(router *gin.Engine) {
	//set up swagger to see all the methods in the handlers
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	router.POST("/user/register", h.Register)
	router.PUT("/user/update/:id", h.UpdateUser)
	router.POST("/user/login", h.LoginUser)
	router.POST("/user/logout", h.Logout)
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/templates/*")
	router.Static("/static", "./static")
}
