package database

import (
	"github.com/lutfiharidha/sequis-test/app"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	// db.AutoMigrate(&app.User{})
	db.AutoMigrate(&app.Friend{})
}
