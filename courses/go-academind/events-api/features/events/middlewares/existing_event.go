package middlewares

import (
	"app/features/events/models"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExistingEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventid")

	event, err := models.GetByID(eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Event #%s not found", eventId),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Cannot get event #%s", eventId),
		})
		return
	}

	ctx.Set("event", event)
	ctx.Next()
}
