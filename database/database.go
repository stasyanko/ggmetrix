package database

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/stasyanko/ggmetrix/models"
	"github.com/stasyanko/ggmetrix/utils"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USERNAME") + " dbname=" + os.Getenv("DB_DATABASE") + " password=" + os.Getenv("DB_PASSWORD")
	//TODO: apply here strategy pattern and log>fatal if db not supported
	db, err := gorm.Open("postgres", dbConfig)

	db.AutoMigrate(&models.Data{}, &models.MetricsType{})

	db.LogMode(utils.IsDevEnv())
	if err != nil {
		panic(err)
	}

	return db, err
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
