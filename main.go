package main

import (
	"fmt"
	"html/template"
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

// indexHandler renders the HTML template with todo data
func indexHandler(c *gin.Context) {
	var todos []Todo
	DB.Find(&todos)
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Template error: %s", err)
		return
	}
	tmpl.Execute(c.Writer, todos)
}

// createTodo handles creating a new todo
func createTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.String(http.StatusBadRequest, "Title cannot be empty")
		return
	}

	todo := Todo{Title: title}
	DB.Create(&todo)
	c.Redirect(http.StatusSeeOther, "/")
}

// deleteTodo handles deleting a todo by ID
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	DB.Delete(&Todo{}, id)
	c.Redirect(http.StatusSeeOther, "/")
}

func main() {
	initDB()
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Routes
	r.GET("/", indexHandler)
	r.POST("/create", createTodo)
	r.GET("/delete/:id", deleteTodo)

	fmt.Println("Server is running at http://localhost:8080")
	r.Run(":8080")
}
