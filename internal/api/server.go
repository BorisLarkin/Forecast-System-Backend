package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Forecast struct {
	id          int
	img_url     string
	title       string
	short_title string
	desc        string
	color       string
}

var Forecasts []Forecast = []Forecast{
	{
		id:          1,
		title:       "Прогноз температуры",
		short_title: "Температура",
		desc:        "Предскажем температуру посредством применения метода авторегрессии",
		color:       "FFC3B6",
		img_url:     "http://127.0.0.1:9000/test/source_obj/temp.png",
	},
	{
		id:          2,
		title:       "Предсказать давление",
		short_title: "Давление",
		desc:        "Покажем в мм рт. ст. наиболее вероятного значения атмосферного давления",
		color:       "D5CEFF",
		img_url:     "http://127.0.0.1:9000/test/source_obj/pressure.png",
	},
	{
		id:          3,
		title:       "Предугадать влажность",
		short_title: "Влажность",
		desc:        "Подскажем, как одеться по влажности атмосферного воздуха, в процентах",
		color:       "DFE5FF",
		img_url:     "http://127.0.0.1:9000/test/source_obj/humidity.png",
	},
}

type Prediction struct {
	id        int
	forecast  Forecast
	date_time string
	place     string
}

var Predictions []Prediction = []Prediction{
	{
		id:        1,
		date_time: "18.09.2024, 19:54",
		place:     "Москва",
		forecast:  Forecasts[0],
	},
	{
		id:        2,
		date_time: "17.09.2024, 14:55",
		place:     "Санкт-Петербург",
		forecast:  Forecasts[2],
	},
	{
		id:        3,
		date_time: "20.10.2024, 00:43",
		place:     "Москва",
		forecast:  Forecasts[1],
	},
}

func StartServer() {

	jsonForecasts, err := os.ReadFile("forecasts.json")
	if err != nil {
		log.Print(err)
	}
	jsonPredictions, err := os.ReadFile("predictions.json")
	json.Unmarshal(jsonForecasts, &Forecasts)
	json.Unmarshal(jsonPredictions, &Predictions)
	if err != nil {
		log.Print(err)
	}

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
			"Forecasts": Predictions,
		})
	})

	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK, "details.tmpl", gin.H{
			"id":          1,
			"title":       "Прогноз температуры",
			"short_title": "Температура",
			"desc":        "Предскажем температуру посредством применения метода авторегрессии",
			"color":       "FFC3B6",
			"img_url":     "http://127.0.0.1:9000/test/source_obj/temp.png",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}
