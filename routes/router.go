package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/controllers"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Define the routes
	r.GET("/", controllers.IndexHandler)
	r.POST("/create", controllers.CreateTodo)
	r.GET("/delete/:id", controllers.DeleteTodo)

	return r
}
