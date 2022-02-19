package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * If statements can contain simple statements preceding conditions
 * Variables declared in simple statements are available in all if branches
 * There is no ternary operator in Go
 */
func main() {

	i1, err := strconv.Atoi("69")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i1)
	}

	// Alternative: Simple if statement to the rescue
	// err is scoped here, there is no conflict with the existing err variable
	if i2, err := strconv.Atoi("42"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i2)
	}

	// Convert kilometers to miles
	// go run main.go 123 // Prints "123 km = 76.383000 mi"
	if args := os.Args; len(args) != 2 {
		fmt.Println("Error: Missing argument")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Error: Argument must be an integer")
	} else {
		miles := float64(km) * 0.621
		fmt.Printf("%d km = %f mi\n", km, miles)
	}
}
