package events

import (
	"app/features/events/handlers"
	eventsMiddlewares "app/features/events/middlewares"
	usersMiddlewares "app/features/users/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	events := routes.Group("/events")
	events.GET("/", handlers.GetEvents)
	events.GET("/:eventid", handlers.GetEvent)

	auth := events.Group("/", usersMiddlewares.Authenticate)
	auth.POST("/", handlers.CreateEvent)

	auth.PUT(
		"/:eventid",
		eventsMiddlewares.ExistingEvent,
		eventsMiddlewares.IsEventAuthor,
		handlers.UpdateEvent,
	)

	auth.DELETE(
		"/:eventid",
		eventsMiddlewares.ExistingEvent,
		eventsMiddlewares.IsEventAuthor,
		handlers.DeleteEvent,
	)
}
