package database

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/stasyanko/ggmetrix/models"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USERNAME") + " dbname=" + os.Getenv("DB_DATABASE") + " password=" + os.Getenv("DB_PASSWORD")
	//TODO: apply here strategy pattern and log>fatal if db not supported
	db, err := gorm.Open("postgres", dbConfig)

	db.AutoMigrate(&models.Data{}, &models.MetricsType{})

	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	return db, err
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
