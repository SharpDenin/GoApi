package config

import (
	"github.com/SharpDenin/PrBack/models"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func connectToDB() {
	db, err := gorm.Open(sqlite.Open("PrBack/config/database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Не подключено")
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
