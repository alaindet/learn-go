package core

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
