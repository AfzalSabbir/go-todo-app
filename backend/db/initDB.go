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
	// Set up the database connection
	DB, err = gorm.Open(sqlite.Open("/mnt/d/DB/SQLite/go-todo-app/app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Perform migration and handle errors
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
