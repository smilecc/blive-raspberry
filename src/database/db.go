package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type SysConfig struct {
	Model
	Name  string `json:"name"`
	Value string `json:"value"`
}

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("blive.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	_ = db.AutoMigrate(&SysConfig{})
	neteaseConfig := &SysConfig{
		Name:  "netease_api_host",
		Value: "https://netease-cloud-music-api-ochre-one.vercel.app",
	}
	db.Where("name = 'netease_api_host'").FirstOrCreate(neteaseConfig)

	DB = db
}
