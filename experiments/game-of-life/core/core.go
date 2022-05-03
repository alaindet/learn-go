package core

import "strings"

type State []bool

type Symbols struct {
	Alive rune
	Dead  rune
}

type Game struct {
	Width      int
	Height     int
	Size       int
	Generation int
	State      []bool
	Symbols
}

func NewGameOfLife(w, h int) *Game {
	size := w * h
	state := make([]bool, size)
	return &Game{
		Width:  w,
		Height: h,
		State:  state,
		Size:   size,
		Symbols: Symbols{
			Alive: '*',
			Dead:  ' ',
		},
	}
}

func (g *Game) SetRandomState(percentageAlive float64) {
	aliveCount := int(percentageAlive * float64(g.Size))
	newState := make([]bool, g.Size)

	// TODO ...
	_ = aliveCount

	g.State = newState
}

func (g *Game) SetSymbols(alive, dead rune) {
	g.Symbols = Symbols{alive, dead}
}

func (g *Game) SetState(s []bool) {
	g.State = s
}

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
