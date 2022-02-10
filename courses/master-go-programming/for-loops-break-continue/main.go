package main

import "fmt"

func main() {

	// continue
	// Skips the current loop iteration and jumps to the next one
	// Ex.: Print even numbers < 10
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// break
	// Interrupts the surrounding loop
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}
		if count == 10 {
			fmt.Println("Maximum count reached")
			break
		}
	}
	fmt.Print("The end") // break exists the loop and jumps here
}
