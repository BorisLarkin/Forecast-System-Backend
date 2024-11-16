package handler

import (
	"web/internal/config"
	"web/internal/ds"
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
	router.DELETE("/forecast/delete/:id", h.WithAuthCheck(ds.Admin), h.DeleteForecast)
	router.POST("/forecast/add", h.WithAuthCheck(ds.Admin), h.AddForecast) //without img
	router.PUT("/forecast/edit/:id", h.WithAuthCheck(ds.Admin), h.EditForecast)
	router.POST("/forecast_to_pred/:forecast_id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.AddForecastToPred)
	router.POST("/forecast/:id/add_picture", h.WithAuthCheck(ds.Admin), h.AddPicture)

	router.GET("/predictions", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.GetPredictions)                //status&form_data filter, no deleted or drafts
	router.GET("/prediction/:id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.GetPredictionById)          //+forecs
	router.DELETE("/prediction/delete/:id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.DeletePrediction) //form_data
	router.PUT("/prediction/form/:id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.FormPrediction)        //client-side
	router.PUT("/prediction/edit/:id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.EditPrediction)        //fields
	router.PUT("/prediction/finish/:id", h.WithAuthCheck(ds.Admin, ds.Moderator), h.FinishPrediction)             //decline or confirm + calc

	router.DELETE("/pr_fc/remove/:prediction_id/:forecast_id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.DeleteForecastFromPred)
	router.PUT("/pr_fc/edit/:prediction_id/:forecast_id", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.EditPredForec)

	router.POST("/user/register", h.Register)
	router.PUT("/user/update/:id", h.WithAuthCheck(ds.Admin), h.UpdateUser)
	router.POST("/user/login", h.LoginUser)                                                 //can proceed anyway
	router.POST("/user/logout", h.WithAuthCheck(ds.Admin, ds.Moderator, ds.User), h.Logout) //have to have any role
}

func (h *Handler) RegisterStatic(router *gin.Engine) {
	router.LoadHTMLGlob("static/templates/*")
	router.Static("/static", "./static")
}
