package main

import (
	"fmt"
	"os"
	"web/internal/config"
	"web/internal/dsn"
	"web/internal/handler"
	"web/internal/pkg"
	"web/internal/repository"

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
	os.Setenv("CURRENT_SESSION", "1")

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
