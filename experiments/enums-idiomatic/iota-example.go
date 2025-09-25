package main

import "fmt"

type LEDColor int

const (
	LEDColorUnknown LEDColor = iota // 0 (Prevents valid zero values)
	LEDColorRed                     // 1
	LEDColorGreen                   // 2
	LEDColorBlue                    // 3
)

func (c LEDColor) String() string {
	switch c {
	case LEDColorRed:
		return "red"
	case LEDColorBlue:
		return "blue"
	case LEDColorGreen:
		return "green"
	default:
		return ""
	}
}

/**
 * Pros
 * - Compile-time constants
 * - Type-safe
 * - Easy to write
 * - Idiomatic in Go (using iota)
 * - Fastest enum-like implementation
 * - Forces explicit type conversion to use integers as the enum
 *
 * Cons
 * - Does not allow string values
 * - Instances have their type in the name to allow namespacing
 * - Values are simply integers, they can't have methods/attributes
 */
func iotaEnumExample() {
	fmt.Println("\n# Iota enum example")

	// Since "myColor" is declared via a constant, it's of "LEDColor" type
	myColor := LEDColorRed

	// This calls String() implicitly
	fmt.Println("myColor is", myColor) // Prints "myColor is red"

	// This maps to LEDColorRed, but it's of type int, not LEDColor
	myOtherColor := 1

	// So, this does not even compile!
	// fmt.Println("myColor is equal to myOtherColor?", myColor == myOtherColor)

	// But this works though
	myThirdColor := LEDColor(myOtherColor)
	fmt.Printf("Same colors? %t\n", myColor == myThirdColor)
}
