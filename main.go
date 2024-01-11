package main

import (
	"example.com/todo-app/constants"
	"example.com/todo-app/db"
	"example.com/todo-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.SetUpDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":" + constants.PORT) // localhost:8080
}
