package main

import (
	"fmt"
	"web/internal/app/config"
	"web/internal/app/dsn"
	"web/internal/app/handler"
	"web/internal/app/repository"
	"web/internal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	router := gin.Default()
	conf, err := config.NewConfig(logger)

	if err != nil {
		logger.Fatalf("Error with configuration reading: #{err}")
	}
	postgresString, errPost := dsn.FromEnv()

	if errPost != nil {
		logger.Fatalf("Error with reading postgres line: #{err}")
	}
	fmt.Println(postgresString)

	rep, errRep := repository.New(postgresString, logger)
	if errRep != nil {
		logger.Fatalf("Error from repo: #{err}")
	}

	hand := handler.NewHandler(logger, rep)
	application := pkg.NewApp(conf, router, logger, hand)
	application.RunApp()
}
