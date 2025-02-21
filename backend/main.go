package main

import (
	"fmt"
	"go-todo-app/db"
	"go-todo-app/routes"
	"log"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()

	fmt.Println("Server is running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
