package main

import (
	"fmt"

	"game_of_life/core"
	"game_of_life/utils"
)

func main() {
	g := core.NewGameOfLife(10, 10)
	g.SetSymbols('*', ' ')
	g.SetState(utils.IntsToBools([]int{
		0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}))
	generations := 20

	printState := func(g *core.Game) {
		fmt.Printf("\nGeneration: %d\n", g.Generation)
		fmt.Println(g.String())
	}

	for i := 0; i <= generations; i++ {
		printState(g)
		g.Step()
	}

	printState(g)
}
