package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &Game{}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
