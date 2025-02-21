package controllers

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/db"
	"go-todo-app/models"
	"html/template"
	"net/http"
)

// IndexHandler renders the HTML template with todo data
func IndexHandler(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Template error: %s", err)
		return
	}
	tmpl.Execute(c.Writer, todos)
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
