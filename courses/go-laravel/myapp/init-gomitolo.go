package main

import (
	"log"
	"os"

	"github.com/alaindet/gomitolo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	g := &gomitolo.Gomitolo{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}

	g.AppName = "myapp"
	g.Debug = true

	app := &application{
		App: g,
	}

	return app
}
