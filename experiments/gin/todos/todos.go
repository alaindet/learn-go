package todos

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a new todo",
	})
}

func GetAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all todos",
	})
}

func GetTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	message := fmt.Sprintf("Get todo #%s", todoID)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func UpdateTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	message := fmt.Sprintf("Update todo #%s", todoID)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	message := fmt.Sprintf("Delete todo #%s", todoID)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
