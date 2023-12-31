package database

import (
	"fmt"
	"log"
	"pomo/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(appConfig *config.AppConfig) error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", appConfig.DBHost, appConfig.DBUserName, appConfig.DBUserPassword, appConfig.DBName, appConfig.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection to Database Failed")
		return err
	}
	fmt.Println("Connection to Database Successfully")
	return nil

}
