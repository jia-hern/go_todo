package routes

import (
	"example.com/todo-app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/todos", getTodos)
	server.GET("/todos/:id", getTodo)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/todos", createTodo)
	authenticated.PUT("/todos/:id", updateTodo)
	authenticated.DELETE("/todos/:id", deleteTodo)
	authenticated.POST("/todos/:id/link", linkTodoToUser)
	authenticated.DELETE("/todos/:id/unlink", unlinkTodoFromUser)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
