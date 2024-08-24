package handlers

import (
	"app/features/events/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventid")

	event, err := models.GetByID(eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Event #%s not found", eventId),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot get event #%s", eventId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Get event #%s", eventId),
		"data":    event,
	})
}
