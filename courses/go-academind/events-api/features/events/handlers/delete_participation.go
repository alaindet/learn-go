package handlers

import (
	"app/common/utils"
	"app/features/events/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEventParticipation(ctx *gin.Context) {

	// Fetch event from middleware
	event, _ := utils.GetFromGinContext[models.EventModel](ctx, "event")
	userId := ctx.GetInt64("userId")

	participation, err := event.DeleteParticipation(userId)
	if err != nil {
		if errors.Is(err, models.ErrEventParticipationMissing) {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("No participation to event #%d exists", event.ID),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not delete your participation to event #%d", event.ID),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "You participation to this event was canceled",
		"data":    participation,
	})
}
