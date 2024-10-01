package controllers

import (
	//"html/template"
	"net/http"
	"web/internal/models"

	"github.com/gin-gonic/gin"
)

func Parse_menu(c *gin.Context) {
	c.HTML(http.StatusOK, "menu.tmpl", gin.H{
		"Forecasts": models.Forecasts,
	})
}
