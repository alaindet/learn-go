package handlers

import (
	"app/common/utils"
	"app/features/events/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvent(ctx *gin.Context) {
	// Fetch event from middleware
	event, _ := utils.GetFromGinContext[models.EventModel](ctx, "event")

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Get event #%s", event.ID),
		"data":    event,
	})
}
