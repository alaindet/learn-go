package main

import (
	"fmt"
	"time"
)

/**
 * Switch is equivalent to a series of if/if-else statements
 * GO strives for simplicity (e.g. no ternary, no while), but switch is an exception
 * as it allows for simpler and shorter code most of the times
 *
 * break statements at the end of each case are not needed
 */
func main() {
	action := "JUMP"

	switch action {

	case "punch": // Single condition
		fmt.Println("Punch!")

	case "jump", "JUMP": // Multiple conditions
		fmt.Println("Jump!")

	case "crouch":
		fmt.Println("Crouch!")

	default: // No condition met
		fmt.Println("No action")
	}

	n := 5

	// You can skip the condition of the switch by either using "switch true"
	// or simply "switch" without a condition

	//switch true {
	switch {

	// Expressions can be more than constants
	case n%2 == 0:
		fmt.Printf("%d is even\n", n)
	default:
		// case n%2 != 0:
		fmt.Printf("%d is odd\n", n)
	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("It's morning")
	case hour < 17:
		fmt.Println("It's afternoon")
	default:
		fmt.Println("It's evening")
	}
}
