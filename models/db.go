package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var Db *gorm.DB

func InitDB() {
	var err error
	Db, err = gorm.Open("postgres", "host=localhost port=5432 user=wemp dbname=wemp password=wemp sslmode=disable")

	if err != nil {
		log.Panic(err)
	}

	Db.AutoMigrate(
		&Domain{},
		&Ip{},
	)
}