package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventid")

	event, err := fetchEvent(ctx, eventId)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Get event #%s", eventId),
		"data":    event,
	})
}
