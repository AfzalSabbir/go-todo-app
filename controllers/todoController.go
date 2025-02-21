package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/db"
	"go-todo-app/models"
	"net/http"
)

// GetTodos retrieves all todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// GetTodo retrieves a single todo by ID
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"todo": todo})
}

// UpdateTodo updates an existing todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.DB.Save(&todo)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
