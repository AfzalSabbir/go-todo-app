package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/db"
	"go-todo-app/models"
	"net/http"
)

// IndexHandler renders the HTML template with todo data
func IndexHandler(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)

	// Load the layout and page templates
	c.HTML(http.StatusOK, "/pages/list.html", gin.H{
		"Title": "Todo List",
		"Todos": todos,
	})
}

// CreateHandler renders the HTML template with todo data
func CreateHandler(c *gin.Context) {
	// Load the layout and page templates
	c.HTML(http.StatusOK, "/pages/add.html", gin.H{
		"Title": "Create Todo",
	})
}

// DetailsHandler renders the HTML template with todo data
func DetailsHandler(c *gin.Context) {
	id := c.Param("id")

	// Declare a variable to hold the todo
	var todo models.Todo

	// Fetch the todo record by its primary key; using First is more appropriate for a single record
	if err := db.DB.First(&todo, id).Error; err != nil {
		// If not found or an error occurs, return a 404
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Load the layout and page templates with the fetched todo data
	c.HTML(http.StatusOK, "/pages/details.html", gin.H{
		"Title": "Todo Details",
		"Todo":  todo,
	})
}

// DetailsHandler renders the HTML template with todo data
func EditHandler(c *gin.Context) {
	id := c.Param("id")

	// Declare a variable to hold the todo
	var todo models.Todo

	// Fetch the todo record by its primary key; using First is more appropriate for a single record
	if err := db.DB.First(&todo, id).Error; err != nil {
		// If not found or an error occurs, return a 404
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Load the layout and page templates with the fetched todo data
	c.HTML(http.StatusOK, "/pages/edit.html", gin.H{
		"Title": "Todo Edit",
		"Todo":  todo,
	})
}

// CreateTodo handles creating a new todo
func CreateTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.String(http.StatusBadRequest, "Title cannot be empty")
		return
	}

	todo := models.Todo{Title: title}
	db.DB.Create(&todo)
	c.Redirect(http.StatusSeeOther, "/")
}

// DeleteTodo handles deleting a todo by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	db.DB.Delete(&models.Todo{}, id)
	c.Redirect(http.StatusSeeOther, "/")
}
