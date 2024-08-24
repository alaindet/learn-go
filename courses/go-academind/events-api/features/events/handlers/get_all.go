package handlers

import (
	"app/features/events/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {

	events, err := models.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot get events",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all events",
		"data":    events,
	})
}
