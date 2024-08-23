package handlers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get events",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Events",
		"data":    events,
	})
}
