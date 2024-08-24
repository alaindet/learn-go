package events

import (
	"app/features/events/handlers"
	eventsMiddlewares "app/features/events/middlewares"
	usersMiddlewares "app/features/users/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	routes.GET("/events", handlers.GetEvents)

	routes.GET(
		"/events/:eventid",
		eventsMiddlewares.ExistingEvent,
		handlers.GetEvent,
	)

	routes.POST(
		"/events",
		usersMiddlewares.Authenticate,
		handlers.CreateEvent,
	)

	routes.PUT(
		"/events/:eventid",
		usersMiddlewares.Authenticate,
		eventsMiddlewares.ExistingEvent,
		eventsMiddlewares.IsEventAuthor,
		handlers.UpdateEvent,
	)

	routes.DELETE(
		"/events/:eventid",
		usersMiddlewares.Authenticate,
		eventsMiddlewares.ExistingEvent,
		eventsMiddlewares.IsEventAuthor,
		handlers.DeleteEvent,
	)

	routes.POST(
		"/events/:eventid/participation",
		usersMiddlewares.Authenticate,
		eventsMiddlewares.ExistingEvent,
		handlers.CreateEventParticipation,
	)

	routes.DELETE(
		"/events/:eventid/participation",
		usersMiddlewares.Authenticate,
		eventsMiddlewares.ExistingEvent,
		handlers.DeleteEventParticipation,
	)
}
