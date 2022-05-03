package main

import (
	"fmt"

	"game_of_life/core"
	"game_of_life/utils"
)

func main() {
	g := core.NewGameOfLife(5, 5)
	g.SetSymbols('x', 'o')
	g.SetState(utils.IntsToBools([]int{
		0, 0, 0, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 0, 0, 0,
	}))

	for i := 0; i < 5; i++ {
		g.Step()
		state := g.String()
		fmt.Printf("\nGeneration: %d\n", g.Generation)
		fmt.Println(state)
	}
}
