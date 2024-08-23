package events

import (
	"app/features/events/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	// TODO: Grouping
	server.POST("/events", handlers.CreateEvent)
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:eventid", handlers.GetEvent)
	server.PUT("/events/:eventid", handlers.UpdateEvent)
	server.DELETE("/events/:eventid", handlers.DeleteEvent)
}
