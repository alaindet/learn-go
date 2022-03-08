package main

import "math"

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
