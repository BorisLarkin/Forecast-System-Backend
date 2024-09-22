package api

import (
	"html/template"
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type Forecast struct {
	Id      int
	Img_url string
	Title   string
	Short   string
	Desc    string
	Color   string
}

var Forecasts []Forecast = []Forecast{
	{
		Id:      1,
		Title:   "Прогноз температуры",
		Short:   "Температура",
		Desc:    "Предскажем температуру посредством применения метода авторегрессии",
		Color:   "(255, 195, 182, 1)",
		Img_url: "http://127.0.0.1:9000/test/source_obj/temp.png",
	},
	{
		Id:      2,
		Title:   "Предсказать давление",
		Short:   "Давление",
		Desc:    "Покажем в мм рт. ст. наиболее вероятного значения атмосферного давления",
		Color:   "(213, 206, 255, 1)",
		Img_url: "http://127.0.0.1:9000/test/source_obj/pressure.png",
	},
	{
		Id:      3,
		Title:   "Предугадать влажность",
		Short:   "Влажность",
		Desc:    "Подскажем, как одеться по влажности атмосферного воздуха, в процентах",
		Color:   "(223, 229, 255, 1)",
		Img_url: "http://127.0.0.1:9000/test/source_obj/humidity.png",
	},
}

type Prediction struct {
	Id        int
	F_id      int
	Date_time string
	Place     string
}

var Predictions []Prediction = []Prediction{
	{
		Id:        1,
		Date_time: "18.09.2024, 19:54",
		Place:     "Москва",
		F_id:      1,
	},
	{
		Id:        2,
		Date_time: "17.09.2024, 14:55",
		Place:     "Санкт-Петербург",
		F_id:      3,
	},
	{
		Id:        3,
		Date_time: "20.10.2024, 00:43",
		Place:     "Москва",
		F_id:      2,
	},
}

type Forecast_parse struct {
	Forecast    Forecast
	Solid_style template.CSS
	Fade_style  template.CSS
}
type Prediction_parse struct {
	Prediction  Prediction
	Forecast    Forecast
	Solid_style template.CSS
	Fade_style  template.CSS
}

var Forecast_parses []Forecast_parse
var Prediction_parses []Prediction_parse

func StartServer() {

	log.Println("Server start up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/menu", func(c *gin.Context) {
		Forecast_parses = Forecast_parses[:0]
		for i, v := range Forecasts {
			var Forecast_parse_tmp Forecast_parse
			Forecast_parse_tmp.Forecast = v
			Forecast_parse_tmp.Solid_style = template.CSS("background-color: rgba" + Forecasts[i].Color)
			Forecast_parse_tmp.Fade_style = template.CSS("background-image:linear-gradient(0deg, rgba" + Forecasts[i].Color + " 0%,rgba(255, 255, 255, 0) 100%)")
			Forecast_parses = append(Forecast_parses, Forecast_parse_tmp)
		}

		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"Forecasts": Forecast_parses,
		})
	})

	r.GET("/cart", func(c *gin.Context) {
		Prediction_parses = Prediction_parses[:0]
		for i, v := range Predictions {
			var Prediction_parse_tmp Prediction_parse
			Prediction_parse_tmp.Prediction = Predictions[i]
			Prediction_parse_tmp.Forecast = Forecasts[slices.IndexFunc(Forecasts, func(f Forecast) bool { return f.Id == v.F_id })]
			var color = Prediction_parse_tmp.Forecast.Color
			Prediction_parse_tmp.Solid_style = template.CSS("background-color: rgba" + color)
			Prediction_parse_tmp.Fade_style = template.CSS("background-image:linear-gradient(0deg, rgba" + Forecasts[i].Color + " 0%,rgba(255, 255, 255, 0) 100%)")
			Prediction_parses = append(Prediction_parses, Prediction_parse_tmp)
		}
		c.HTML(http.StatusOK, "cart.tmpl", gin.H{
			"Predictions": Prediction_parses,
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK, "details.tmpl", gin.H{
			"Id":      1,
			"Title":   "Прогноз температуры",
			"Short":   "Температура",
			"Desc":    "Предскажем температуру посредством применения метода авторегрессии",
			"img_url": "http://127.0.0.1:9000/test/source_obj/temp.png",
		})
	})
	r.Static("/assets", "./resources")
	//r.Static("/favicon.ico", "./resources/source_obj/favicon.ico")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}
