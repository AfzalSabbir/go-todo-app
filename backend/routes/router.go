package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/controllers"
)

// SetupRouter sets up the routes for the API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define API routes
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	return r
}
