package main

import (
	"fmt"
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
type MetricsTypeModel = models.MetricsType
type CounterRequest struct {
	Title string `form:"title" binding:"required"`
}

func incrementCounter(title string) {
	counterLock.Lock()
	defer counterLock.Unlock()
	if _, ok := counters[title]; ok {
		counters[title] = counters[title] + 1
	} else {
		// create new counter in metrics_types
		newCounterType := MetricsTypeModel{
			Title: title,
			Type:  "counter",
		}
		db.NewRecord(newCounterType)
		err := db.Create(&newCounterType).Error
		// TODO: refactore to return error instead of panicking??
		if err != nil {
			panic(err)
		}
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
	var countersFromDB []MetricsTypeModel
	if err := db.Select("DISTINCT(title), type").Where("type = ?", "counter").Find(&countersFromDB).Error; err != nil {
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
				err := db.Create(&newCounter).Error
				if err != nil {
					fmt.Println(err)
				}
				counters[counterTitle] = 0
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
	router.GET("/counter/:title", func(c *gin.Context) {
		title := c.Params.ByName("title")
		var counterData []DataModel
		fromTime := int32(time.Now().Unix()) - 86400

		if err := db.Select("unix_ts, value, title").Where("title = ?", title).Where(" type = ?", "counter").Where("unix_ts >= ?", fromTime).Find(&counterData).Error; err != nil {
			c.AbortWithStatus(400)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": counterData})
		}
	})
	// increment counter
	router.POST("/counter", func(c *gin.Context) {
		var counterRequest CounterRequest
		c.BindJSON(&counterRequest)
		incrementCounter(counterRequest.Title)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Counter created"})
	})
	router.GET("/select_options", func(c *gin.Context) {
		var metricsTypes []MetricsTypeModel

		if err := db.Find(&metricsTypes).Error; err != nil {
			c.AbortWithStatus(400)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": metricsTypes})
		}
	})

	// Start server
	router.Run(":8000")
}

func createCounter(title string) {
	counterLock.Lock()
	defer counterLock.Unlock()
	counters[title] = 0
}
