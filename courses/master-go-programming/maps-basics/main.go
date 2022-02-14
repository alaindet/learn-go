package main

import (
	"fmt"
)

func mapsBasics() {
	var employees map[string]string
	fmt.Printf("Type: %T, Value: %#v\n", employees, employees)
	// Type: map[string]string, Value: map[string]string(nil)
	fmt.Printf("Number of pairs: %d\n", len(employees)) // Number of pairs: 0

	// Accessing a non-existing key returns the zero value of the map
	fmt.Printf("Value of \"John\" is %q\n", employees["John"]) // Value of "John" is ""

	var accounts map[string]int
	fmt.Printf("%#v\n", accounts["fortytwo"]) // 0 (zero value of int is 0)

	// Error: Slices are not comparable
	// var m1 map[[]int]string // Error: invalid map key type

	// Arrays are comparable and can be used as keys!
	var m1 map[[2]int]string
	_ = m1

	// Error: Cannot set a key of a nil map, initialize it first!
	// employees["Dan"] = "Programmer"

	// Initialize an empty map
	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 28
	fmt.Println(people) // map[John:21.4 Mary:28]

	// Alternative: initialize an empty map
	m2 := make(map[int]int)
	m2[69] = 42
	fmt.Println(m2) // map[69:42]

	// Initialize map with a multi-line literal
	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13, // <- trailing comma is required on each line!
	}
	_ = balances

	// Initialize map with an inline literal (no trailing comma required)
	m3 := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m3

	// When setting a key-value pair, you either override or create a key-value pair
	balances["USD"] = 500.2
	balances["GBP"] = 100.3

	// When reading a key-value pair, you either get the real value or a zero-value as default
	// The "comma ok" idiom is used to check for existence of a key-value pair specifically
	// The second value returned is a boolean indicating existence of key-value pair
	value, ok := balances["RON"]
	fmt.Println(value, ok) // 0 false
	value2, ok := balances["USD"]
	fmt.Println(value2, ok) // 500.2 true

	// Looping over a map
	// NOTE: THIS IS DISCOURAGE as it's slow and maps are not ordered
	for key, value := range balances {
		fmt.Printf("%s => %f\n", key, value)
		// USD => 500.200000
		// EUR => 432.130000
		// GBP => 100.300000
	}
}

func mapsComparison() {
	a := map[string]string{"A": "B"}
	b := map[string]string{"A": "B"}
	// areEqual := a == b // Invalid operation

	// Compare string representations of maps
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	fmt.Println(s1, s2)                      // map[A:B] map[A:B]
	fmt.Println("Are maps equal?", s1 == s2) // Are maps equal? true

	// Note:
	// The string representation of a map has keys sorted so that
	// two represenations of two maps with the same key-value pairs are equal!
	c := map[string]string{"a": "b", "c": "d"}
	d := map[string]string{"c": "d", "a": "b"}
	s3 := fmt.Sprintf("%s", c)
	s4 := fmt.Sprintf("%s", d)
	fmt.Println(s3, s4)                      // map[a:b c:d] map[a:b c:d]
	fmt.Println("Are maps equal?", s3 == s4) // Are maps equal? true
}

func mapsCloning() {
	friends := map[string]int{
		"Alice":   10,
		"Bob":     20,
		"Charlie": 30,
	}
	// This is just a new reference of the same underlying map header
	somePeople := friends
	// This affects somePeople as well
	friends["Alice"] = 20
	fmt.Println(somePeople) // map[Alice:20 Bob:20 Charlie:30]

	// "Clone" a map
	somePeopleClone := make(map[string]int)
	for key, value := range somePeople {
		somePeopleClone[key] = value
	}

	somePeople["Bob"] = 25                 // Overwrite key-value pair
	delete(somePeople, "Charlie")          // Delete key-value pair
	delete(somePeople, "non-existing-key") // This is fine, does not throw error

	fmt.Println(somePeople) // map[Alice:20 Bob:25]
	// The clone was not modified!
	fmt.Println(somePeopleClone) //  map[Alice:20 Bob:20 Charlie:30]
}

func main() {
	mapsBasics()
	mapsComparison()
	mapsCloning()
}
