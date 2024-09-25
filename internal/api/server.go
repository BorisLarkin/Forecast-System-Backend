package api

import (
	"html/template"
	"log"
	"net/http"
	"slices"
	"strconv"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	log.Println("Server start up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/menu", func(c *gin.Context) {
		models.Forecast_parses = models.Forecast_parses[:0]
		for i, v := range models.Forecasts {
			var Forecast_parse_tmp models.Forecast_parse
			Forecast_parse_tmp.Forecast = v
			Forecast_parse_tmp.Solid_style = template.CSS("background-color: rgba" + models.Forecasts[i].Color)
			Forecast_parse_tmp.Fade_style = template.CSS("background-image:linear-gradient(0deg, rgba" + models.Forecasts[i].Color + " 0%,rgba(255, 255, 255, 0) 100%)")
			models.Forecast_parses = append(models.Forecast_parses, Forecast_parse_tmp)
		}

		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"Forecasts": models.Forecast_parses,
		})
	})

	r.GET("/cart", func(c *gin.Context) {
		models.Prediction_parses = models.Prediction_parses[:0]
		for i, v := range models.Predictions {
			var Prediction_parse_tmp models.Prediction_parse
			Prediction_parse_tmp.Prediction = models.Predictions[i]
			Prediction_parse_tmp.Forecast = models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == v.F_id })]
			var color = Prediction_parse_tmp.Forecast.Color
			Prediction_parse_tmp.Solid_cart_style = template.CSS("background-color: rgba" + color)
			Prediction_parse_tmp.Fade_cart_style = template.CSS("background-image:linear-gradient(90deg, rgba" + color + " 0%,rgba(255, 255, 255, 0) 100%)")
			models.Prediction_parses = append(models.Prediction_parses, Prediction_parse_tmp)
		}
		c.HTML(http.StatusOK, "cart.tmpl", gin.H{
			"Predictions": models.Prediction_parses,
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	r.GET("/details", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		Forecast_to_det := models.Forecasts[slices.IndexFunc(models.Forecasts, func(f models.Forecast) bool { return f.Id == id })]
		c.HTML(http.StatusOK, "details.tmpl", gin.H{
			"Forecast_to_det": Forecast_to_det,
		})
	})
	r.Static("/assets", "./resources")
	//r.Static("/favicon.ico", "./resources/source_obj/favicon.ico")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}
