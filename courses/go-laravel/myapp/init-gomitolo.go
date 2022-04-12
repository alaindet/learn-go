package main

import (
	"log"
	"myapp/handlers"
	"os"

	"github.com/alaindet/gomitolo"
)

func initApplication() *application {

	// Get current path
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Init new app
	g := &gomitolo.Gomitolo{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	// Setup app
	g.AppName = os.Getenv("APP_NAME")
	appHandlers := &handlers.Handlers{
		App: g,
	}

	app := &application{
		App:      g,
		Handlers: appHandlers,
	}

	// Add routes
	app.App.Routes = app.routes()

	return app
}
