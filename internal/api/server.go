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

	r.GET("/menu", controllers.Parse_menu)

	r.GET("/cart", controllers.ParseCart)

	/*
		r.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.tmpl", gin.H{})
		})
	*/

	r.GET("/details", controllers.Details_parse)

	r.Static("/assets", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
