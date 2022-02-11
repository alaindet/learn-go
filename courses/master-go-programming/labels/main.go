package main

import (
	"fmt"
)

/**
 * There is only one this to know about labels:
 * PLEASE AVOID LABELS AND GOTO AT ANY COST
 *
 * Labels can identify lines of code in order to refer to them
 * - They can be used with break, continue and goto
 * - They MUST be used once declared
 * - They do not conflict with any variable name and are not block scoped
 * - They are scoped inside the function in which they are declared
 * - They are rarely used and generally discouraged, mostly used for exiting outer loops
 * - You cannot jump to a label if a new variable would be created
 */
func main() {
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Mary"}
	outer := "I do not conflict with the outer label"
	_ = outer

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}

	fmt.Println("I am past outer label")

	i := 0
onceMore:
	if i < 5 {
		fmt.Println("Another round? Hell yeah", i)
		i++
		goto onceMore
	}
	fmt.Println("I am past onceMore label")

	// This is not possible since you would skip a variable declaration
	// (And you would skip the line...)
	// 	goto cutTheLine
	// 	peopleInLine := 42
	// cutTheLine:
	// 	fmt.Println("I cut the line!")
}
