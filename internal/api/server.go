package api

import (
	"log"
	"web/internal/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	log.Println("Server start up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/forecasts", controllers.Parse_forecasts) //1

	r.GET("/prediction/:id", controllers.Parseprediction) //2

	r.GET("/details", controllers.Details_parse) //3
	//4 - POST del pred; 5 - POST del forec
	r.Static("/assets", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
