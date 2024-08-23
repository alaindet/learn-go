package main

import (
	"app/core/db"
	"app/features/events"
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
	events.Routes(server)

	server.Run(":" + cfg.Port)
}
