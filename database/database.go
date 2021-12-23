package database

import (
	"log"
	"os"

	"github.com/dkantikorn/fiber-rest-api/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var Database DatabaseInstance

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running for daatbase migration")

	// migration models
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DatabaseInstance{DB: db}
}
