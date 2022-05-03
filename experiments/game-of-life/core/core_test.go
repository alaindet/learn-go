package core

import (
	"testing"

	"game_of_life/utils"
)

func TestGameOfLife(t *testing.T) {
	t.Run("basic representationof 5x5 grid", func(t *testing.T) {
		g := NewGameOfLife(5, 5)
		g.SetState(utils.IntsToBools([]int{
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
		}))

		result := g.String()
		expected := utils.JoinLines(
			"     ",
			"  *  ",
			"  *  ",
			"  *  ",
			"     ",
		)

		if result != expected {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})

	t.Run("change symbols", func(t *testing.T) {
		g := NewGameOfLife(3, 3)
		g.SetSymbols('@', '-')
		g.SetState(utils.IntsToBools([]int{
			0, 1, 0,
			0, 1, 1,
			0, 1, 0,
		}))

		result := g.String()
		expected := utils.JoinLines(
			"-@-",
			"-@@",
			"-@-",
		)

		if result != expected {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})

	t.Run("calculate 1 step", func(t *testing.T) {
		g := NewGameOfLife(5, 5)
		g.SetSymbols('x', 'o')
		g.SetState(utils.IntsToBools([]int{
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
		}))
		g.Step()

		result := g.String()
		expected := utils.JoinLines(
			"ooooo",
			"ooooo",
			"oxxxo",
			"ooooo",
			"ooooo",
		)

		if result != expected {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})

	t.Run("calculate multiple steps", func(t *testing.T) {
		g := NewGameOfLife(5, 5)
		g.SetSymbols('x', 'o')
		g.SetState(utils.IntsToBools([]int{
			0, 0, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 0, 0,
		}))
		g.Steps(4)

		result := g.String()
		expected := utils.JoinLines(
			"ooooo",
			"ooxoo",
			"ooxoo",
			"ooxoo",
			"ooooo",
		)

		if result != expected {
			t.Errorf("Result: %v Expected: %v", result, expected)
		}
	})
}
