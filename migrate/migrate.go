package main

import (
	"fmt"
	"log"
	"pomo/config"

	"pomo/internal/database"
	"pomo/internal/models"

	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.Connection(config)
	databaseMigrations(database.DB)
}

func databaseMigrations(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(&models.UserModel{})

	fmt.Println("Database migrations successfully")

}
