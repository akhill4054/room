package models

import (
	"fmt"

	"github.com/akhill4054/room-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	db = _db

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Password{})
	db.AutoMigrate(&UserProfile{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&RoomMember{})
	// db.AutoMigrate(&models.Message{})
}
