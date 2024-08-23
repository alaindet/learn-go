package main

import (
	"app/db"
	"app/handlers"
	"flag"

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

	server.POST("/events", handlers.CreateEvent)
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:eventid", handlers.GetEvent)

	server.Run(":" + cfg.Port)
}
