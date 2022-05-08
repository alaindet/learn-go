package core

import (
	"fmt"
	"strings"
	"time"
)

func (g *Game) String() string {
	var b strings.Builder

	for i := 0; i < g.Size; i++ {

		if g.State[i] {
			b.WriteRune(g.Symbols.Alive)
		} else {
			b.WriteRune(g.Symbols.Dead)
		}

		if (i+1)%g.Width == 0 && i+1 != g.Size {
			b.WriteRune('\n')
		}
	}

	return b.String()
}

func (g *Game) Print() {
	fmt.Println("Game of Life")
	fmt.Printf("Generation: %d\n", g.Generation)
	fmt.Printf("Grid: %d x %d (Size: %d)\n", g.Width, g.Height, g.Size)
	fmt.Printf("Alive cells: %q, Dead cells: %q\n", g.Alive, g.Dead)
	fmt.Printf("======\n\n%s\n\n", g.String())
}

func (g *Game) Animate(generations int, duration time.Duration) {

	print := func() {
		// https://imzye.com/Go/golang-clear-screen/
		fmt.Printf("\x1bc") // Clear console on each frame
		fmt.Print(g.String())
	}

	for i := 0; i <= generations; i++ {
		print()
		g.Step()
		time.Sleep(duration)
	}

	print()
	fmt.Printf("\nDone\n")
}
