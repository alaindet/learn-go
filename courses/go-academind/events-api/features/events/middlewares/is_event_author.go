package middlewares

import (
	"app/common/utils"
	"app/features/events/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsEventAuthor(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")

	event, err := utils.GetFromGinContext[models.EventModel](ctx, "event")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if event.UserID != userId {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You must be the author of the event to update or delete it",
		})
		return
	}

	ctx.Set("isAuthor", true)
	ctx.Next()
}
