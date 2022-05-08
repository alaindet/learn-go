package core

import (
	"math/rand"
	"time"
)

func (g *Game) SetSize(sizes ...int) {
	w, h := 0, 0

	// Square grid
	if len(sizes) == 1 {
		w, h = sizes[0], sizes[0]
	}

	// Rectangle grid, ignore any other input
	if len(sizes) > 1 {
		w, h = sizes[0], sizes[1]
	}

	g.Width = w
	g.Height = h
	g.Size = w * h
	g.State = make([]bool, g.Size)
}

func (g *Game) SetRandomState(percentageAlive float64) {
	newState := make([]bool, g.Size)
	randomizer := getRandomizer()
	threshold := percentageAlive * 100.0

	for i := 0; i < g.Size; i++ {
		newState[i] = float64(randomizer.Intn(100)) < threshold
	}

	g.State = newState
}

func (g *Game) SetSymbols(alive, dead rune) {
	g.Symbols = Symbols{alive, dead}
}

func (g *Game) SetState(s []bool) {
	g.State = s
}

func getRandomizer(args ...int64) *rand.Rand {
	seed := time.Now().UTC().UnixMilli()
	if len(args) > 0 {
		seed = args[0]
	}
	randomSource := rand.NewSource(seed)
	return rand.New(randomSource)
}
