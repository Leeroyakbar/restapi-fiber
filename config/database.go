package config

import (
	"github.com/leeroyakbar/restapi-fiber/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DSN"),
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(
		model.Book{},
	)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	DB = db
}
