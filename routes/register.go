package routes

import (
	"net/http"
	"strconv"

	"example.com/todo-app/models"
	"github.com/gin-gonic/gin"
)

func linkTodoToUser(context *gin.Context) {
	userId := context.GetInt64("userId")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id."})
		return
	}

	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo."})
		return
	}

	err = todo.LinkTodoToUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not link user to todo."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Linked user to todo."})
}

func unlinkTodoFromUser(context *gin.Context) {
	userId := context.GetInt64("userId")
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id."})
		return
	}

	var todo models.Todo
	todo.Id = todoId

	err = todo.UnlinkTodoFromUser(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unlink user from todo."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Unlinked user from todo."})
}
