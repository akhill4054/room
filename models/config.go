package models

import (
	"fmt"

	"github.com/akhill4054/room-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDB() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_HOST,
		config.POSTGRES_DB,
	)
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db = _db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Password{})
	// db.AutoMigrate(&UserProfile{})
	// db.AutoMigrate(&models.Room{})
	// db.AutoMigrate(&models.Message{})
}
