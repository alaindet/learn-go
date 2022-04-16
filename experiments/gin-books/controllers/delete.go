package controllers

import (
	"fmt"
	"gin_books/app"
	"gin_books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DELETE /books/:bookid
func DeleteBook(c *gin.Context) {
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

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book #%d deleted", book.ID),
		"data":    book,
	})
}
