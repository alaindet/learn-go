/////////////////////////////////
// Named (Defined) Types in Go
// Go Playground: https://play.golang.org/p/v2-QZsESmC-
/////////////////////////////////

package main

import "fmt"

type age int        //new type, int is the underlying type
type oldAge age     //new type, int (not age) is the underlying type
type veryOldAge age //new type, int (not age) is the underlying type

func example() {

	// new type speed (underlying type uint)
	type speed uint

	// s1, s2 of type speed
	var s1 speed = 10
	var s2 speed = 20

	// performing operations with the new types
	fmt.Println(s2 - s1) // -> 10

	// uint and speed are different types (they have different names)
	var x uint

	// x = s1  //error different types

	// correct
	x = uint(s1)
	_ = x

	// correct
	var s3 speed = speed(x)
	_ = s3

	// declaring a variable of type uint8
	var a uint8 = 10
	var b byte // byte is an alias to uit8

	// even though they have different names, byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	// type alias_name = type_name
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hou
}
