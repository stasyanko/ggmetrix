package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"

	//TODO: fork this repo and replace its link to fork
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"

	"github.com/stasyanko/ggmetrix/database"
)

var db *gorm.DB

//TODO: use struct for tasks and tasks are to be renamed to smth else
var tasks []string

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

	tasks = append(tasks, "1")
	tasks = append(tasks, "2")
	tasks = append(tasks, "3")

	cronObj := cron.New()
	cronObj.AddFunc("0 * * * * *", func() {
		//TODO: init tasks from DB in a loop
		//TODO: new tasks will be in tasks slice
		// just one AddFunc is enough, all tasks are run
		// in a loop, each of them in its own goroutine
		for _, task := range tasks {
			dt := time.Now()
			fmt.Println("Test save to DB at: "+task, dt.String())
		}
	})
	cronObj.Start()

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
