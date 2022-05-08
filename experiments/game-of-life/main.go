package main

import (
	"time"

	// "game_of_life/core"
	"game_of_life/examples"
)

func main() {
	n := 200
	total := time.Second * 10
	symbols := [2]rune{'*', ' '}

	// examples.RunExample(examples.Glider, n, total, symbols)
	examples.RunExample(examples.GosperGliderGun, n, total, symbols)
	// examples.RunExample(examples.Spaceships, n, total, symbols)
	// examples.RunExample(examples.Pulsar, n, total, symbols)
}

// func randomExample() {
// 	n := 1000
// 	total := time.Second * 10

// 	g := core.NewGameOfLife(50, 50)
// 	g.SetRandomState(0.66)
// 	g.Animate(n, total)
// }
