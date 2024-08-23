package handlers

import (
	"app/features/events/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetchEvent(ctx *gin.Context, eventId string) (models.EventModel, error) {
	event, err := models.GetByID(eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Event #%s not found", eventId),
			})
			return models.EventModel{}, err
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot get event #%s", eventId),
		})
		return models.EventModel{}, err
	}

	return event, nil
}
