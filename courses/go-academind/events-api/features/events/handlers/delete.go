package handlers

import (
	"app/common/utils"
	"app/features/events/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEvent(ctx *gin.Context) {

	// Fetch event from middleware
	event, _ := utils.GetFromGinContext[models.EventModel](ctx, "event")

	// Delete model
	deletedEvent, err := event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot delete event on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Event #%d deleted", event.ID),
		"data":    deletedEvent,
	})
}
