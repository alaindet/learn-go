package main

import (
	"fmt"
	"gin-example/app"
	"gin-example/todos"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin example")
	app := app.NewApp()

	r := gin.Default()

	r.GET("/todos", todos.GetAllTodos)
	r.GET("/todos/:todoID", todos.GetTodo)
	r.POST("/todos", todos.CreateTodo)
	r.PUT("/todos/:todoID", todos.UpdateTodo)
	r.DELETE("/todos/:todoID", todos.DeleteTodo)

	r.GET("/lab/querystring", queryParamsExample)
	r.POST("/lab/body", requestBodyExample)
	r.PUT("/lab/body", requestBodyExample)
	r.PATCH("/lab/body", requestBodyExample)

	addr := fmt.Sprintf(":%s", app.Config.Port)
	r.Run(addr)
}
