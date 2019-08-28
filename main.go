package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"

	//TODO: fork this repo and replace its link to fork
	"github.com/foolin/gin-template/supports/gorice"
	"github.com/gin-gonic/gin"

	"github.com/stasyanko/ggmetrix/database"
	"github.com/stasyanko/ggmetrix/models"
)

var db *gorm.DB

//counters for counting the number of smth in mterics
var counters = make(map[string]Counter)
var counterLock = sync.RWMutex{}

type Counter int
type DataModel = models.Data
type CounterRequest struct {
	Title string `form:"title" binding:"required"`
}

func incrementCounter(title string) {
	counterLock.Lock()
	defer counterLock.Unlock()
	if _, ok := counters[title]; ok {
		counters[title] = counters[title] + 1
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

	// init counters
	// TODO: move to separate func
	var countersFromDB []DataModel
	if err := db.Select("DISTINCT(title)").Where("type = ?", "counter").Find(&countersFromDB).Error; err != nil {
		panic(err)
	}
	for _, counter := range countersFromDB {
		createCounter(counter.Title)
	}
}

func main() {
	defer db.Close()

	cronObj := cron.New()
	cronObj.AddFunc("0 * * * * *", func() {
		for k := range counters {
			go func(counterTitle string) {
				counterLock.Lock()
				defer counterLock.Unlock()
				newCounter := DataModel{
					Title:  counterTitle,
					Type:   "counter",
					Value:  uint(counters[counterTitle]),
					UnixTs: uint(time.Now().Unix()),
				}
				db.NewRecord(newCounter)
				db.Create(&newCounter)
				counters[counterTitle] = 0
				// fmt.Println("key:", counters[counterTitle])
			}(k)
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
	router.POST("/counter", func(c *gin.Context) {
		var counterRequest CounterRequest
		c.BindJSON(&counterRequest)
		incrementCounter(counterRequest.Title)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Counter created"})
	})

	// Start server
	router.Run(":8000")
}

func createCounter(title string) {
	counterLock.Lock()
	defer counterLock.Unlock()
	counters[title] = 0
}
