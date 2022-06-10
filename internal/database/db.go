package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"vaultea/api/internal/models"
)

var once sync.Once
var db *gorm.DB

func ConnectToDatabase() {
	once.Do(func() {
		dsn := "vaultea:123456@tcp(127.0.0.1:9000)/vaultea?charset=utf8mb4&parseTime=True&loc=Local" // TODO: Make configurable
		connectedDb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		db = connectedDb
	})
}

func MakeMigrations() {
	db.AutoMigrate(&models.Vault{}, &models.User{}, &models.Folder{}, &models.Password{}, &models.UserPreferences{})
}

func GetDb() gorm.DB {
	return *db
}
