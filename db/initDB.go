package db

import (
	"go-todo-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// DB is a global variable for the GORM database connection
var DB *gorm.DB

// InitDB initializes the SQLite database and auto-migrates the todos table
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate will create the table if it doesn't exist
	DB.AutoMigrate(&models.Todo{})
}
