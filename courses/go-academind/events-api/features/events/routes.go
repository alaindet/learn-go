package events

import (
	"app/features/events/handlers"
	"app/features/users/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	events := routes.Group("/events")
	events.GET("/", handlers.GetEvents)
	events.GET("/:eventid", handlers.GetEvent)

	auth := events.Group("/", middlewares.Authenticate)
	auth.POST("/", handlers.CreateEvent)
	auth.PUT("/:eventid", handlers.UpdateEvent)
	auth.DELETE("/:eventid", handlers.DeleteEvent)
}
