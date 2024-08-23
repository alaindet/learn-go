package main

import (
	"app/db"
	"app/handlers"
	"app/models"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port string
}

func main() {
	db.InitDB()
	cfg := ServerConfig{}
	flag.StringVar(&cfg.Port, "port", "8080", "Server port")
	flag.Parse()

	server := gin.Default()

	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handleCreateEvent)

	server.Run(":" + cfg.Port)
}

// TODO: Move
func handleCreateEvent(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create event due to invalid data",
		})
		return
	}

	event.UserID = 1 // TODO
	savedEvent, err := event.Save()
	if err != nil {
		fmt.Println(err.Error()) // TODO: Remove
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create event on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"data":    savedEvent,
	})
}
