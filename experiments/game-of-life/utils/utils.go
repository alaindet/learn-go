package utils

import (
	"fmt"
	"strings"

	"game_of_life/core"
)

func IntsToBools(ints []int) []bool {
	result := make([]bool, len(ints))
	for i, val := range ints {
		if val == 0 {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}

func JoinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}

func PrintGameGenerations(g *core.Game, generations int) {
	for i := 0; i <= generations; i++ {
		PrintGameState(g)
		g.Step()
	}

	PrintGameState(g)
}

func PrintGameState(g *core.Game) {
	fmt.Printf("\nGeneration: %d\n", g.Generation)
	fmt.Println(g.String())
}
