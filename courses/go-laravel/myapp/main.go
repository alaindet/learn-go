package main

import (
	"github.com/alaindet/gomitolo"
)

type application struct {
	App *gomitolo.Gomitolo
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
