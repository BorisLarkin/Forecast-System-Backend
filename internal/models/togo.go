package models

import "html/template"

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
		Img_url: "http://127.0.0.1:9000/test/source_obj/temp_crop.png",
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
	Result    string
}

var Predictions []Prediction = []Prediction{
	{
		Id:        1,
		Date_time: "18.09.2024, 19:54",
		Result:    "760 мм. рт. ст.",
		F_id:      1,
	},
	{
		Id:        2,
		Date_time: "17.09.2024, 14:55",
		Result:    "38%",
		F_id:      3,
	},
	{
		Id:        3,
		Date_time: "20.10.2024, 00:43",
		Result:    "В работе...",
		F_id:      2,
	},
}

type Forecast_parse struct {
	Forecast    Forecast
	Solid_style template.CSS
	Fade_style  template.CSS
}
type Prediction_parse struct {
	Prediction       Prediction
	Forecast         Forecast
	Solid_cart_style template.CSS
	Fade_cart_style  template.CSS
}

var Forecast_parses []Forecast_parse
var Prediction_parses []Prediction_parse

var HeaderDiv template.HTML = template.HTML(`
	<div class=header_component>
      <div class="header_bg"></div>
      <div class=logo>
        <button class="logo_btn" onclick="location.href='http://127.0.0.1:8080/menu'"></button>
        <span  class="logo_lbl">Погода</span>
        <div class="logo_img"></div>
      </div>
    </div>`)
