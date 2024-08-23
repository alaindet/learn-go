package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEvent(ctx *gin.Context) {

	// Parse route param
	eventId := ctx.Param("eventid")
	event, err := fetchEvent(ctx, eventId)
	if err != nil {
		return
	}

	// Delete model
	deletedEvent, err := event.Delete()
	if err != nil {

		// TODO: Remove
		fmt.Printf("%#v\n", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot delete event on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Event #%s deleted", eventId),
		"data":    deletedEvent,
	})
}
