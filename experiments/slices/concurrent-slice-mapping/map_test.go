package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestMapDummySource struct {
	ID string
}

type TestMapDummyDest struct {
	ID int
}

func TestMap(t *testing.T) {

	multiplyByTen := func(n int) int {
		return n * 10
	}

	mapTestDummySourceToDest := func(source TestMapDummySource) TestMapDummyDest {
		id, _ := strconv.Atoi(source.ID)
		return TestMapDummyDest{ID: id}
	}

	t.Run("maps a slice of integers", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6}
		expected := []int{10, 20, 30, 40, 50, 60}
		outcome := Map(elements, multiplyByTen)
		assert.Equal(t, outcome, expected)
	})

	t.Run("maps a slice of strings", func(t *testing.T) {
		elements := []string{"foo", "bar", "baz"}
		expected := []string{"FOO", "BAR", "BAZ"}
		outcome := Map(elements, strings.ToUpper)
		assert.Equal(t, outcome, expected)
	})

	t.Run("maps a slice of structs", func(t *testing.T) {
		elements := []TestMapDummySource{{"11"}, {"22"}, {"33"}}
		expected := []TestMapDummyDest{{11}, {22}, {33}}
		outcome := Map(elements, mapTestDummySourceToDest)
		assert.Equal(t, outcome, expected)
	})
}
