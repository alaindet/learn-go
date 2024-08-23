package handlers

import (
	"app/features/events/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateEvent(ctx *gin.Context) {

	// Parse JSON body
	var updatedEvent models.EventModel
	err := ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse input data",
		})
	}

	// Parse route param
	eventId := ctx.Param("eventid")
	event, err := fetchEvent(ctx, eventId)
	if err != nil {
		return
	}

	// Update existing model
	event.Name = updatedEvent.Name
	event.Description = updatedEvent.Description
	event.Location = updatedEvent.Location
	event.DateTime = updatedEvent.DateTime

	// Save updated model
	savedEvent, err := event.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot update event on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Event #%s updated", eventId),
		"data":    savedEvent,
	})
}
