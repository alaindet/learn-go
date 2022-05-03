package core

func (g *Game) Step() {
	size := len(g.State)
	newState := make([]bool, size)

	for i := 0; i < size; i++ {
		neighbors := g.CountNeighbors(i)
		newState[i] = false // Cells die unless specific conditions arise
		isAlive := g.State[i]
		shouldSurvive := isAlive && (neighbors == 2 || neighbors == 3)
		shouldBeBorn := !isAlive && neighbors == 3
		if shouldSurvive || shouldBeBorn {
			newState[i] = true
		}
	}

	g.Generation++
	g.State = newState
}

func (g *Game) Steps(steps int) {
	for i := 0; i < steps; i++ {
		g.Step()
	}
}

func (g *Game) CountNeighbors(i int) int {
	neighbors := 0
	size := len(g.State)
	lookArounds := []int{
		i - g.Width,
		i - g.Width + 1,
		i + 1,
		i + g.Width + 1,
		i + g.Width,
		i + g.Width - 1,
		i - 1,
		i - g.Width - 1,
	}

	for _, lookAround := range lookArounds {

		// No neighbor in this direction, it's considered dead
		if lookAround < 0 || lookAround >= size {
			continue
		}

		if g.State[lookAround] {
			neighbors++
		}
	}

	return neighbors
}
