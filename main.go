package main

import (
	"log"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	//TODO: fork this repo and replace its link to fork
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"

	"github.com/stasyanko/ggmetrix/database"
)

var db *gorm.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	var dbErr error
	db, dbErr = database.Initialize()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}

func main() {
	defer db.Close()

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
