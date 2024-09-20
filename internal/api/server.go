package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Forecast struct {
	Id          int
	Img_url     string
	Title       string
	Short_Title string
	Desc        string
	Color       string
}

var Forecasts []Forecast = []Forecast{
	{
		Id:          1,
		Title:       "Прогноз температуры",
		Short_Title: "Температура",
		Desc:        "Предскажем температуру посредством применения метода авторегрессии",
		Color:       "(255, 195, 182, 1)",
		Img_url:     "http://127.0.0.1:9000/test/source_obj/temp.png",
	},
	{
		Id:          2,
		Title:       "Предсказать давление",
		Short_Title: "Давление",
		Desc:        "Покажем в мм рт. ст. наиболее вероятного значения атмосферного давления",
		Color:       "(213, 206, 255, 1)",
		Img_url:     "http://127.0.0.1:9000/test/source_obj/pressure.png",
	},
	{
		Id:          3,
		Title:       "Предугадать влажность",
		Short_Title: "Влажность",
		Desc:        "Подскажем, как одеться по влажности атмосферного воздуха, в процентах",
		Color:       "(223, 229, 255, 1)",
		Img_url:     "http://127.0.0.1:9000/test/source_obj/humidity.png",
	},
}

type Prediction struct {
	Id        int
	Forecast  Forecast
	Date_time string
	Place     string
}

var Predictions []Prediction = []Prediction{
	{
		Id:        1,
		Date_time: "18.09.2024, 19:54",
		Place:     "Москва",
		Forecast:  Forecasts[0],
	},
	{
		Id:        2,
		Date_time: "17.09.2024, 14:55",
		Place:     "Санкт-Петербург",
		Forecast:  Forecasts[2],
	},
	{
		Id:        3,
		Date_time: "20.10.2024, 00:43",
		Place:     "Москва",
		Forecast:  Forecasts[1],
	},
}

func StartServer() {

	log.Println("Server start up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/menu", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"Forecasts": Forecasts,
		})
	})

	r.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.tmpl", gin.H{
			"Predictions": Predictions,
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK, "details.tmpl", gin.H{
			"Id":          1,
			"Title":       "Прогноз температуры",
			"Short_Title": "Температура",
			"Desc":        "Предскажем температуру посредством применения метода авторегрессии",
			"img_url":     "http://127.0.0.1:9000/test/source_obj/temp.png",
		})
	})
	r.Static("/assets", "./resources")
	//r.Static("/favicon.ico", "./resources/source_obj/favicon.ico")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}
