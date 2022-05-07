package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type SysConfig struct {
	gorm.Model
	Name  string
	Value string
}

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("blive.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	_ = db.AutoMigrate(&SysConfig{})
	DB = db
}
