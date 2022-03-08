package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func mathBasics() {
	val1 := 279.00
	val2 := 48.95

	p("Abs: %v", math.Abs(val1))                 // Abs: 279
	p("Ceil: %v", math.Ceil(val2))               // Ceil: 49
	p("Copysign: %v", math.Copysign(val1, -5))   // Copysign: -279
	p("Floor: %v", math.Floor(val2))             // Floor: 48
	p("Max: %v", math.Max(val1, val2))           // Max: 279
	p("Min: %v", math.Min(val1, val2))           // Min: 48.95
	p("Mod: %v", math.Mod(val1, val2))           // Mod: 34.249999999999986
	p("Pow: %v", math.Pow(val1, 2))              // Pow: 77841
	p("Round: %v", math.Round(val2))             // Round: 49
	p("RoundToEven: %v", math.RoundToEven(48.5)) // RoundToEven: 48
}

func getRandomInteger(from, to int) int {
	return from + rand.Intn(to-from+1)
}

func randomBasics() {
	// WARNING: This prints the same values every time since it has the same seed
	for i := 0; i < 5; i++ {
		p("Static random %v : %v", i, rand.Int())
	}

	fmt.Println(strings.Repeat("#", 20))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		p("Dynamic random %v : %v", i, rand.Int())
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < 10; i++ {
		p("Random from 10 to 20: %d", getRandomInteger(10, 20))
	}
}

func shuffling() {
	var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

	getSwapFn := func(elements []string) func(i, j int) {
		return func(i, j int) {
			elements[i], elements[j] = elements[j], elements[i]
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), getSwapFn(names))

	for i, name := range names {
		p("Index %v: Name: %v", i, name)
	}
}
