package main

import (
	"fmt"
)

func main() {
	s1 := "Hello World"
	fmt.Printf("%q\n", s1) // "Hello World"

	// Enclosing quotes using escaped double quotes
	s2 := "He said \"OK\""
	_ = s2

	// Enclosing quotes using raw string
	// Note: a "raw string" is a string enclosed in backticks
	s3 := `She said "Fine"`
	_ = s3

	// These two are equivalent
	fmt.Println("This is\nAwesome")
	fmt.Println(`This is
Awesome`)

	// No escaping for raw strings
	fmt.Println(`C:\Users\You`)

	// Concatenation
	s4 := "I am learning the " + "Go " + "programming language"
	fmt.Println(s4 + "!") // I am learning the Go programming language!

	// Indexing
	char1 := s4[0]
	fmt.Println(char1) // 73 <-- ASCII value for 'I'

	// Run time error: cannot assign a new letter to a string as it's immutable
	// s4[0] = 'Y'
}
