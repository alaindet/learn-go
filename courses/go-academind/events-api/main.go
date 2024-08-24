package main

import (
	"app/core/db"
	"app/core/server"
	"app/features/events"
	"app/features/users"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set tu
	db.InitDB()
	cfg := server.ReadConfigFromCLI()

	// Routes
	server := gin.Default()
	routes := server.Group("/api/v1")
	users.Routes(routes)
	events.Routes(routes)

	// Print all registered routes
	if gin.IsDebugging() {
		debugRoutes(server)
	}

	// Bootstrap
	err := server.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}

func debugRoutes(server *gin.Engine) {

	routes := make([]string, 0)

	for _, r := range server.Routes() {
		route := fmt.Sprintf("%s %s", r.Method, r.Path)
		routes = append(routes, route)
	}

	fmt.Printf(
		"\nDEBUG REGISTERED ROUTES\n=======================\n\n%s\n\n",
		strings.Join(routes, "\n"),
	)
}
