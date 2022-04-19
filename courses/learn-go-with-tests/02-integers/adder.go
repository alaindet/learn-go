package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	r := Add(40, 2)
	fmt.Println("Result is", r)
}
