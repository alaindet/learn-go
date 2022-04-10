package main

import (
	"fmt"

	"github.com/alaindet/gomitolo"
)

type application struct {
	App *gomitolo.Gomitolo
}

func main() {
	initApplication()
	fmt.Println("Hello World")
}
