package controllers

import (
	"fmt"
	"gin_books/app"
	"gin_books/dto"
	"gin_books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PATCH /books/:bookid
func UpdateBook(c *gin.Context) {
	db := app.GetApp().Database
	bookID := c.Param("id")
	var dtoIn dto.UpdateBookDto
	var book models.Book

	err := c.ShouldBindJSON(&dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.Where("id = ?", bookID).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Book #%s not found", bookID),
		})
		return
	}

	db.Model(&book).Updates(dtoIn)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book #%d update", book.ID),
		"data":    book,
	})
}
