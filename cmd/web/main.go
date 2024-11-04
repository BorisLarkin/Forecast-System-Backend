package main

import (
	"fmt"
	"web/internal/config"
	"web/internal/dsn"
	"web/internal/handler"
	"web/internal/minio"
	"web/internal/pkg"
	"web/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	router := gin.Default()
	router.Use(CORSMiddleware())
	conf, err := config.NewConfig(logger)
	minioClient := minio.NewMinioClient(conf)

	if err != nil {
		logger.Fatalf("Error with configuration reading: #{err}")
	}
	postgresString, errPost := dsn.FromEnv()

	if errPost != nil {
		logger.Fatalf("Error with reading postgres line: #{err}")
	}
	fmt.Println(postgresString)

	rep, errRep := repository.New(postgresString, logger, minioClient)
	if errRep != nil {
		logger.Fatalf("Error from repo: #{err}")
	}

	hand := handler.NewHandler(logger, rep, minioClient)
	application := pkg.NewApp(conf, router, logger, hand)
	application.RunApp()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
