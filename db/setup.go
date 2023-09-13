package db

import (
	"fmt"
	"os"

	"crispypod.com/crispypod/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	if len(dbPort) == 0 {
		dbPort = "5432"
	}
	var dbConnString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err))
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create User table: %s", err))
	}

	err = db.AutoMigrate(&models.Episode{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create Episode table: %s", err))
	}

	err = db.AutoMigrate(&models.SiteConfig{})
	if err != nil {
		panic(fmt.Sprintf("Failed to create SiteConfig table: %s", err))
	}

	DB = db
}
