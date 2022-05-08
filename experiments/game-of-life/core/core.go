package core

import (
	"strings"
)

func NewGameOfLife(sides ...int) *Game {
	g := &Game{
		Width:  0,
		Height: 0,
		State:  make([]bool, 0),
		Size:   0,
		Symbols: Symbols{
			Alive: '*',
			Dead:  ' ',
		},
	}

	g.SetSize(sides...)

	return g
}

func FromString(alive, dead rune, grid string) *Game {

	lines := strings.Split(grid, "\n")
	w := len([]rune(lines[0]))
	h := len(lines)
	state := make([]bool, w*h)

	for i, line := range lines {
		for j, char := range []rune(line) {
			pos := i*w + j
			state[pos] = char == alive
		}
	}

	return &Game{
		Width:  w,
		Height: h,
		State:  state,
		Size:   w * h,
		Symbols: Symbols{
			Alive: alive,
			Dead:  dead,
		},
	}
}
