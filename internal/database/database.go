package database

import (
	"fmt"
	"pmsys/internal/app/models"
	"pmsys/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectAndMigrate() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Env("DB_USER", "root"),
		config.Env("DB_PASS", ""),
		config.Env("DB_HOST", "127.0.0.1"),
		config.Env("DB_PORT", "3306"),
		config.Env("DB_NAME", "app_db"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Category{}); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
