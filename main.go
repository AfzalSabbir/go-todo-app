package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Todo represents a task in the todo list
type Todo struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null"`
}

// DB is a global variable for the GORM database connection
var DB *gorm.DB

// initDB initializes the SQLite database and auto-migrates the todos table
func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate will create the table if it doesn't exist
	DB.AutoMigrate(&Todo{})
}

// getTodos retrieves all todos and returns them as JSON
func getTodos(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

// createTodo handles creating a new todo
func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

// deleteTodo handles deleting a todo by ID
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	DB.Delete(&Todo{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}

func main() {
	initDB()
	r := gin.Default()

	// Routes
	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.DELETE("/todos/:id", deleteTodo)

	fmt.Println("Server is running at http://localhost:8080")
	r.Run(":8080")
}
