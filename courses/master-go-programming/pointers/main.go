package main

import (
	"fmt"
)

func main() {
	name := "John"
	addr := &name
	fmt.Println(name, addr) // John 0xc000010230

	var x int = 2
	ptr := &x
	// %p prints a pointer address in hexadecimal notation
	fmt.Printf("Type of pointer: %T\n", ptr)        // Type of pointer: *int
	fmt.Printf("Address of x: %p\n", ptr)           // Address of x: 0xc000014080
	fmt.Printf("Address of ptr itself: %p\n", &ptr) // Address of ptr itself: 0xc00000e030

	// Uninitialized pointer
	var myPointer *int
	fmt.Printf("%v\n", myPointer) // <nil>

	// Uninitialized pointer (alternative)
	p := new(int)
	x = 100
	p = &x                    // Assign address of x to p
	fmt.Printf("%T %v", p, p) // *int 0xc000014080
	*p = 90                   // equivalent to x = 90
	fmt.Println(x)            // 90
	fmt.Println(*p == x)      // true

	// Be aware of how to use the star symbol *!
	aaa := 1
	// Here, it means p will reference a variable of type int
	var ppp *int = &aaa
	// Here, you are dereferencing p to use x's value directly
	bbb := *ppp
	fmt.Println(aaa, bbb, ppp) // 1 1 0xc000014090

	myVal := 5.5
	myPtr1 := &myVal
	myPtr2 := &myPtr1

	// val: 0xc00012a020, addr: 0xc000122028
	fmt.Printf("val: %v, addr: %v\n", myPtr1, &myPtr1)

	// val: 0xc000122028, addr: 0xc000122030
	fmt.Printf("val: %v, addr: %v\n", myPtr2, &myPtr2)

	fmt.Println(
		*myPtr1,             // 5.5 The value of myVal
		*myPtr2,             // <hex> The value of myPtr1
		**myPtr2,            // 5.5 The value of myVal
		*myPtr1 == **myPtr2, // true
	)

	**myPtr2++
	fmt.Println(myVal) // 6.5

	var p1 *int
	var p2 *int
	fmt.Println(p1 == p2) // true (Both are <nil>)
	myVal2 := 123
	myVal3 := 123
	p1 = &myVal2
	p2 = &myVal2
	fmt.Println(p1 == p2) // true (Both point to the same variable)
	p2 = &myVal3
	fmt.Println(p1 == p2) // false (Because myVal2 and myVal3 have different addresses!)
}
