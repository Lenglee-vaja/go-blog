package database

import (
	"log"
	"os"

	"github.com/lenglee-vaja/blogbackend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cloud not connect to database")
	} else {
		log.Println("Connected to database")
	}
	DB = db

	db.AutoMigrate(&model.User{})
}
