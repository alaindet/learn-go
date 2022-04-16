package controllers

import (
	"fmt"
	"gin_books/app"
	"gin_books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
func FindBooks(c *gin.Context) {
	db := app.GetApp().Database

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"message": "All books data",
		"data":    books,
	})
}

// GET /books/:bookid
func FindBook(c *gin.Context) { // Get model if exist
	db := app.GetApp().Database
	bookID := c.Param("id")

	var book models.Book
	err := db.Where("id = ?", bookID).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Book #%s not found", bookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book #%d bound", book.ID),
		"data":    book,
	})
}
