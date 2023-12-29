package main

import (
	"fmt"
	"gin_books/app"
	"gin_books/controllers"
	"gin_books/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.Setup()
	router := gin.Default()
	routes(router)
	bootstrap(router)
}

func bootstrap(router *gin.Engine) {
	address := fmt.Sprintf(":%s", app.Get("APP_PORT"))
	router.Run(address) // :8080 default
}

func routes(router *gin.Engine) {
	// TODO: Group routes?
	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	// Add routes here...
}
