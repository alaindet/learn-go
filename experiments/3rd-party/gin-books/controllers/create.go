package controllers

import (
	"gin_books/app"
	"gin_books/dto"
	"gin_books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /books
func CreateBook(c *gin.Context) {
	db := app.GetApp().Database
	var dtoIn dto.CreateBookDto

	err := c.ShouldBindJSON(&dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:  dtoIn.Title,
		Author: dtoIn.Author,
	}

	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created",
		"data":    book,
	})
}
