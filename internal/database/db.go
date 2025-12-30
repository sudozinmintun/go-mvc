package database

import (
	"go-pongo2-demo/internal/app/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectAndMigrate() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.User{}, &models.Category{})

	// seed admin
	var count int64
	DB.Model(&models.User{}).Count(&count)
	if count == 0 {
		models.SeedAdmin(DB)
	}
}
