package handlers

import (
	"app/common/utils"
	"app/features/events/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEventParticipation(ctx *gin.Context) {

	// Fetch event from middleware
	event, _ := utils.GetFromGinContext[models.EventModel](ctx, "event")
	userId := ctx.GetInt64("userId")

	participation, err := event.CreateParticipation(userId)
	if err != nil {

		if errors.Is(err, models.ErrEventParticipationExists) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": fmt.Sprintf("You already added your participation for event #%d", event.ID),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not add your participation to event #%d", event.ID),
			"data":    event,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "You participation to this event was saved",
		"data":    participation,
	})
}
