package models

import (
	"html/template"
	"slices"
	"strings"
)

type Forecast struct {
	Id            int
	Img_url       string
	Title         string
	Short         string
	Desc          string
	Color         string
	Measure_type  string
	Extended_desc string
}

var Forecasts []Forecast = []Forecast{
	{
		Id:            0,
		Title:         "Прогноз температуры",
		Short:         "Температура",
		Desc:          "Предскажем температуру посредством применения метода скользящего среднего",
		Color:         "255, 195, 182, 1",
		Img_url:       "http://127.0.0.1:9000/test/source_obj/temp_crop.png",
		Measure_type:  "градусы цельсия, °C",
		Extended_desc: "Нахождение вероятных значений средних температур на последующие дни с учетом тренда изменения средней за скользящее окно дней температуры.",
	},
	{
		Id:           1,
		Title:        "Предсказать давление",
		Short:        "Давление",
		Desc:         "Покажем в мм рт. ст. наиболее вероятного значения атмосферного давления",
		Color:        "213, 206, 255, 1",
		Img_url:      "http://127.0.0.1:9000/test/source_obj/pressure.png",
		Measure_type: "миллиметры ртутного столба, мм рт. ст.",
		Extended_desc: `Нахождение на последующие дни наиболее вероятного диапазона давлений с учетом тренда.
		Следует учитывать вероятность зацикливания предсказаний в следствие ограниченности диапазона давлений.`,
	},
	{
		Id:           2,
		Title:        "Предугадать влажность",
		Short:        "Влажность",
		Desc:         "Подскажем, как одеться по влажности атмосферного воздуха, в процентах",
		Color:        "223, 229, 255, 1",
		Img_url:      "http://127.0.0.1:9000/test/source_obj/humidity.png",
		Measure_type: "проценты, %",
		Extended_desc: `Нахождение статистически вероятных значений влажности на последующие дни с учетом тренда изменения средних значений влажности.
		Стоит отметить, что метод будет работать исправно лишь при использовании больших значений величины скользящего окна:
		в долгосрочной перспективе учитываются макроизменения, которые зачастую вбирают в себя все микроизменения, чего, очевидно, не происходит в короткосрочных измерениях.
		В последних любые экстримальные значения, обусловленные осадками или аномалиями, значительно снизят точность предсказания.`,
	},
}

type Prediction struct {
	Id             int
	Date_created   string
	Date_formed    string
	Date_completed string
	Creator        string
	Moderator      string
	Status         string
}

type Prediction_Forecasts struct {
	Predicition_id int
	Forecast_id    int
	Measures       string
	Result         string
}

var Predictions []Prediction = []Prediction{
	{
		Id:             0,
		Date_created:   "",
		Date_formed:    "",
		Date_completed: "",
		Creator:        "",
		Moderator:      "",
		Status:         "done",
	},
	{
		Id:             1,
		Date_created:   "",
		Date_formed:    "",
		Date_completed: "",
		Creator:        "",
		Moderator:      "",
		Status:         "done",
	},
	{
		Id:             2,
		Date_created:   "",
		Date_formed:    "",
		Date_completed: "",
		Creator:        "",
		Moderator:      "",
		Status:         "in-work",
	},
}
var Prediction_Forecasts_arr []Prediction_Forecasts = []Prediction_Forecasts{
	{
		Predicition_id: 0,
		Forecast_id:    1,
		Measures:       "",
		Result:         "",
	},
	{
		Predicition_id: 0,
		Forecast_id:    2,
		Measures:       "",
		Result:         "",
	},
	{
		Predicition_id: 0,
		Forecast_id:    0,
		Measures:       "",
		Result:         "",
	},
}

var HeaderDiv template.HTML = template.HTML(`
	<div class=header_component>
      <div class="header_bg"></div>
      <div class=logo>
        <button class="logo_btn" onclick="location.href='http://127.0.0.1:8080/forecasts'"></button>
        <span  class="logo_lbl">Погода</span>
        <div class="logo_img"></div>
      </div>
    </div>`)

func GetForecastById(id int) Forecast {
	return Forecasts[slices.IndexFunc(Forecasts, func(f Forecast) bool { return f.Id == id })]
}
func GetPredictionById(id int) Prediction {
	return Predictions[slices.IndexFunc(Predictions, func(f Prediction) bool { return f.Id == id })]
}
func GetForecastsByPredictionId(id int) []Forecast {
	var Fs []Forecast
	for i, v := range Prediction_Forecasts_arr {
		if Prediction_Forecasts_arr[i].Predicition_id == id {
			Fs = append(Fs, GetForecastById(v.Forecast_id))
		}
	}
	return Fs
}

var CURRENT_PREDICTION_ID = 0

func GetPredictionLen(id int) int {
	var len = 0
	for i := range Prediction_Forecasts_arr {
		if Prediction_Forecasts_arr[i].Predicition_id == id {
			len += 1
		}
	}
	return len
}

func GetCartLen() int {
	return GetPredictionLen(CURRENT_PREDICTION_ID)
}
func GetForecasts() []Forecast {
	return Forecasts
}
func GetPredictions() []Prediction {
	return Predictions
}
func GetCurrPredictionId() int {
	return CURRENT_PREDICTION_ID
}
func SetCurrPredictionId(id int) {
	CURRENT_PREDICTION_ID = id
}
func GetForecastsByName(name_like string) []Forecast {
	var Fs []Forecast
	for i := range Forecasts {
		if strings.Contains(Forecasts[i].Title, name_like) || strings.Contains(Forecasts[i].Short, name_like) {
			Fs = append(Fs, Forecasts[i])
		}
	}
	return Fs
}
