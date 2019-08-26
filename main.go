package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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

//counters for counting the number of smth in mterics
var counters = make(map[string]Counter)
var counterLock = sync.RWMutex{}

type Counter int

func incrementCounter(title string) {
	counterLock.Lock()
	defer counterLock.Unlock()
	if _, ok := counters[title]; !ok {
		counters[title]++
	} else {
		counters[title] = 1
	}
}

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

	counters["test1"] = 5
	counters["test2"] = 7

	cronObj := cron.New()
	cronObj.AddFunc("* * * * * *", func() {
		for k := range counters {
			go incrementCounter(k)

			fmt.Println("key:", counters[k])
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
