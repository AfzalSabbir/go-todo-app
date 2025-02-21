package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-todo-app/controllers"
	"go-todo-app/db"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Routes
	r.GET("/", controllers.IndexHandler)
	r.POST("/create", controllers.CreateTodo)
	r.GET("/delete/:id", controllers.DeleteTodo)

	fmt.Println("Server is running at http://localhost:8080")
	r.Run(":8080")
}
