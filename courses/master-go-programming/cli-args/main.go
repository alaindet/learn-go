package main

import (
	"fmt"
	"os"
	"strconv"
)

// Try running this with
// go run main.go hello world
func main() {
	// os.Args is a slice of string arguments passed to the CLI
	// The first argument is always the name of the compiled file just executed
	// fmt.Println("Path:", os.Args[0])
	// fmt.Println("Arg #1:", os.Args[1])
	// fmt.Println("Number of args:", len(os.Args))

	// Ex.: go run main.go 3.14159
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	_ = err
	fmt.Printf("Result (%T) %f\n", result, result) // Result (float64) 3.141590
}
