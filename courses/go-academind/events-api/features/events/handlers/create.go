package handlers

import (
	"app/features/events/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")

	var event models.EventModel
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create event due to invalid data",
		})
		return
	}

	event.UserID = userId
	savedEvent, err := event.Create()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create event on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"data":    savedEvent,
	})
}
