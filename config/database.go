package config

import (
	"github.com/SharpDenin/PrBack/models"
	"gorm.io/gorm"
	// "modernc.org/sqlite"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не подключено")
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
