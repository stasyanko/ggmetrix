package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	//TODO: fork this repo and replace its link to fork
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// servers other static files
	staticBox := rice.MustFindBox("static")
	router.StaticFS("/static", staticBox.HTTPBox())

	//new template engine
	router.HTMLRender = gorice.New(rice.MustFindBox("views"))

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "ggmetrix"})
	})

	router.GET("/chart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chart.html", gin.H{"title": "ggmetrix"})
	})

	// Start server
	router.Run(":8000")
}
