package model

import (
	"log"
	"netcut/config"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open(config.DbPath), &gorm.Config{})
	if err != nil {
		log.Panic("db connect fail:", err)
	}
	db.AutoMigrate(&Cut{})
}
