package routes

import (
	"net/http"
	"strconv"

	"example.com/todo-app/models"
	"github.com/gin-gonic/gin"
)

func getTodos(context *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todos. Try again later."})
		return
	}
	context.JSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id."})
	}

	todo, err := models.GetTodoById(todoId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo."})
		return
	}
	context.JSON(http.StatusOK, todo)
}

func createTodo(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	userId := context.GetInt64("userId")
	todo.UserId = userId

	err = todo.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create todo."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Todo created!", "todo": todo})
}

func updateTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id."})
		return
	}

	userId := context.GetInt64("userId")
	todo, err := models.GetTodoById(todoId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the todo."})
		return
	}

	if todo.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only creator of todo can update the todo."})
		return
	}

	var updatedTodo models.Todo
	err = context.ShouldBindJSON(&updatedTodo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo request data."})
		return
	}

	updatedTodo.Id = todoId
	err = updatedTodo.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update todo."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo update successfully."})
}

func deleteTodo(context *gin.Context) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id."})
		return
	}

	userId := context.GetInt64("userId")
	todo, err := models.GetTodoById(todoId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the todo."})
		return
	}

	if todo.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Only creator of todo can delete the todo."})
		return
	}

	err = todo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully."})
}
