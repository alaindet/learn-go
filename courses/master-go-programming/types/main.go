package main

import "fmt"

func main() {
	var i1 int8 = -128     // Minimum value for int8
	fmt.Printf("%T\n", i1) // int8

	var i2 uint16 = 65535  // Maximum value for uint16
	fmt.Printf("%T\n", i2) // uint16

	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3) // float64 float64 float64

	var myRune rune = 'a'
	var myByte byte = 'b'
	fmt.Printf("%T %T\n", myRune, myByte)
	fmt.Println(myRune, myByte) // 97 98

	var isTrue bool = true
	fmt.Printf("%T\n", isTrue) // bool

	var s1 string = "Hello World"
	fmt.Printf("%T\n", s1) // string

	var nums [4]int = [4]int{4, 5, -1, 10} // Init an array of length 4 with given elements
	fmt.Printf("%T\n", nums)               // [4]int

	var cities []string = []string{"London", "Paris", "Rome"}
	fmt.Printf("%T\n", cities) // []string

	var balances map[string]float64 = map[string]float64{
		"USD": 123.4,
		"EUR": 321.0,
	}
	fmt.Printf("%T\n", balances) // map[string]float64

	type Person struct {
		name string
		age  int
	}

	var me Person
	fmt.Printf("%T\n", me) // main.Person

	var x int = 2
	var myPointer *int = &x
	fmt.Printf("myPointer is of type %T with a value of %v\n", myPointer, myPointer)
	// Prints "myPointer is of type *int with a value of 0xc0000ba028"
	// NOTE: The memory address in hexadecimal format changes on every run

	fmt.Printf("%T\n", myFunction) // func()
}

// Unused function
func myFunction() {}
