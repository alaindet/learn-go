package handlers

import (
	"app/common/utils"
	"app/features/events/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEventParticipations(ctx *gin.Context) {

	// Fetch event from middleware
	event, _ := utils.GetFromGinContext[models.EventModel](ctx, "event")

	users, err := event.GetParticipations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot get participations for event #%d", event.ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Get all participations of event #%d", event.ID),
		"data":    users,
	})
}
