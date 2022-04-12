package main

import (
	"myapp/handlers"

	"github.com/alaindet/gomitolo"
)

type application struct {
	App      *gomitolo.Gomitolo
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
