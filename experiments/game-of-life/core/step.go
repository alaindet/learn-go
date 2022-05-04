package core

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

func (g *Game) CountNeighbors(i int) int {
	neighbors := 0
	lookArounds := []int{
		i - g.Width - 1, // top-left
		i - g.Width,     // top
		i - g.Width + 1, // top-right
		i - 1,           // left
		i + 1,           // right
		i + g.Width - 1, // bottom-left
		i + g.Width,     // bottom
		i + g.Width + 1, // bottom-right
	}

	for _, lookAround := range lookArounds {

		// No neighbor in this direction, it's considered dead
		if lookAround < 0 || lookAround >= len(g.State) {
			continue
		}

		if g.State[lookAround] {
			neighbors++
		}
	}

	return neighbors
}
