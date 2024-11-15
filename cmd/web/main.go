package main

import (
	"fmt"
	"net/http"

	//"web/docs"
	"web/internal/config"
	"web/internal/handler"
	"web/internal/minio"
	"web/internal/pkg"
	redis_api "web/internal/redis-api"
	"web/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title Forecast system
// @version 1.0
// @description Bmstu Open IT Platform

// @contact.name API Support
// @contact.url https://vk.com/b.larkin
// @contact.email borislarkin18@mail.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1
// @schemes https http
// @BasePath /

func main() {
	logger := logrus.New()
	router := gin.Default()
	router.Use(handler.CORSMiddleware())
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Fatal error reading config: %s", err)
	}
	minioClient := minio.NewMinioClient(conf)

	if err != nil {
		logger.Fatalf("Error with configuration reading: #{err}")
	}
	host := conf.Postgresql.DB_Host
	port := conf.Postgresql.DB_Port
	user := conf.Postgresql.DB_User
	pass := conf.Postgresql.DB_Pass
	dbname := conf.Postgresql.DB_Name
	if host == "" || port == "" || user == "" || pass == "" || dbname == "" {
		logger.Fatalf("Database config is broken")
	}
	postgresString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	fmt.Println(postgresString)

	redisClient, err := redis_api.New(conf.Redis)
	if err != nil {
		logger.Fatalf("Error creating redis: #{err}")
	}

	rep, errRep := repository.New(postgresString, logger, minioClient, redisClient)
	if errRep != nil {
		logger.Fatalf("Error from repo: #{err}")
	}

	hand := handler.NewHandler(logger, rep, minioClient, conf)
	application := pkg.NewApp(conf, router, logger, hand)
	application.Router.GET("/ping/:name", Ping)
	application.RunApp()
}

type pingReq struct{}
type pingResp struct {
	Status string `json:"status"`
}

// Ping godoc
// @Summary      Show hello text
// @Description  very very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  pingResp
// @Router       /ping/{name} [get]
func Ping(gCtx *gin.Context) {
	name := gCtx.Param("name")
	gCtx.String(http.StatusOK, "Hello %s", name)
}
