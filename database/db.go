package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
	"log"
	"mygram/models"
)

var (
	host = "localhost"
	user = "postgres"
	password = "admin"
	dbPort = "5432"
	dbname = "mygram"
	db *gorm.DB
	err error
)

func StartDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database")
	db.Debug().AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB{
	return db
}