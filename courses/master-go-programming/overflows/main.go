package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++            // uint8 cannot represent 256, so x silently goes in overflow
	fmt.Println(x) // 0

	// Compile time overflow
	// Overflows are checked only at compile time
	// Examples
	// var y uint8 = 256 // Compiler error: Numeric Overflow
	// a := uint8(255 + 1)

	// Runtime overflow
	var b uint8 = 255
	fmt.Printf("%d\n", b+1) // 0

	// Float numbers overflow to Inf!
	f := float32(math.MaxFloat32)
	fmt.Println(f) // 3.4028235e+38
	f *= 1.1
	fmt.Println(f) // +Inf

	// Compile time error
	// const i int8 = 256
}
