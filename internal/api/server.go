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

var Forecasts []Forecast

type Prediction struct {
	id        int
	f_id      int
	date_time string
	place     string
	f_type    Forecast
}

var Predictions [3]Prediction

func StartServer() {

	jsonForecasts, err := os.ReadFile("forecasts.json")
	jsonPredictions, err := os.ReadFile("predictions.json")
	json.Unmarshal(jsonForecasts, &Forecasts)
	json.Unmarshal(jsonPredictions, &Predictions)
	if err != nil {
		log.Print(err)
	}
	for i, p := range Predictions {
		p.f_type = Forecasts[p.f_id]
		log.Println("got request ", i)
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
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"Forecasts": Forecasts,
		})
	})

	r.GET("/desc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "menu.tmpl", gin.H{
			"Forecasts": Forecasts,
		})
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Println("Server down")
}

/*
package main


import (

	"html/template"

	"log"

	"net/http"

	"os"

	"regexp"

)


type Page struct {

	Title string

	Body  []byte

}


func (p *Page) save() error {

	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)

}


func loadPage(title string) (*Page, error) {

	filename := title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}


func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return

	}

	renderTemplate(w, "view", p)

}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		p = &Page{Title: title}

	}

	renderTemplate(w, "edit", p)

}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}


var templates = template.Must(template.ParseFiles("edit.html", "view.html"))


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}


var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {

			http.NotFound(w, r)

			return

		}

		fn(w, r, m[2])

	}

}


func main() {

	http.HandleFunc("/view/", makeHandler(viewHandler))

	http.HandleFunc("/edit/", makeHandler(editHandler))

	http.HandleFunc("/save/", makeHandler(saveHandler))


	log.Fatal(http.ListenAndServe(":8080", nil))

}
*/
