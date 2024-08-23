package main

import (
	"app/core/db"
	"app/core/server"
	"app/features/events"
	"app/features/users"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set tu
	db.InitDB()
	cfg := server.ReadConfigFromCLI()

	// Routes
	server := gin.Default()
	users.Routes(server)
	events.Routes(server)

	// Bootstrap
	err := server.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
