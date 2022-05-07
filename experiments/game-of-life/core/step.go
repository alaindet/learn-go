package core

// The 8 neighbors positions of a cell in a 2D-grid
const (
	tl = "top-left"
	t  = "top"
	tr = "top-right"
	r  = "right"
	br = "bottom-right"
	b  = "bottom"
	bl = "bottom-left"
	l  = "left"
)

const noNeighbor = -1

func (g *Game) Step() {
	g.Generation++
	newState := make([]bool, len(g.State))

	for i := 0; i < len(g.State); i++ {
		newState[i] = false // Cells die unless specific conditions arise
		n := g.CountNeighbors(i)
		if n == 3 || (n == 2 && g.State[i]) {
			newState[i] = true
		}
	}

	g.State = newState
}

func (g *Game) Steps(steps int) {
	for i := 0; i < steps; i++ {
		g.Step()
	}
}

// Returns the coordinates of all 8 neighbors if possible
// A cylindrical plane is assumed, so right neighbors of a cell on the right edge
// are the ones on the opposite (left) edge and viceversa
// Think about Pac-Man or Snake where you "wrap" on the other side
func (g *Game) getNeighborsPositions(i int) map[string]int {
	n := map[string]int{
		tl: i - g.Width - 1,
		t:  i - g.Width,
		tr: i - g.Width + 1,
		l:  i - 1,
		r:  i + 1,
		bl: i + g.Width - 1,
		b:  i + g.Width,
		br: i + g.Width + 1,
	}

	onTopEdge := n[t] < 0
	onBottomEdge := n[b] > (g.Size - 1)
	onLeftEdge := i%g.Width == 0
	onRightEdge := (i+1)%g.Width == 0

	// Wrap left!
	if onLeftEdge {
		n[tl] = i - 1
		n[l] = i + g.Width - 1
		n[bl] = i + (2 * g.Width) - 1
	}

	// Wrap right!
	if onRightEdge {
		n[tr] = i - (2 * g.Width) + 1
		n[r] = i - g.Width + 1
		n[br] = i + 1
	}

	if onTopEdge {
		n[tl] = noNeighbor
		n[t] = noNeighbor
		n[tr] = noNeighbor
	}

	if onBottomEdge {
		n[bl] = noNeighbor
		n[b] = noNeighbor
		n[br] = noNeighbor
	}

	return n
}

func (g *Game) CountNeighbors(i int) int {
	count := 0
	n := g.getNeighborsPositions(i)

	for _, ni := range n {

		if ni == noNeighbor {
			continue
		}

		if g.State[ni] {
			count++
		}
	}

	return count
}
