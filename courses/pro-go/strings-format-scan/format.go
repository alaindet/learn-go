package main

import (
	"fmt"
)

func stringsFormatting() {

	// Basic printing to standard output
	fmt.Print("foo", "bar", "baz", "\n")

	// Print all arguments in a line
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)

	fmt.Printf("Value is %.2f\n", 3.1415926535) // Value is 3.14

	output := fmt.Sprintf("Value is %.2f", 3.1415926535)
	fmt.Println(output) // Value is 3.14

	err := fmt.Errorf("Value is %.2f\n", 3.1415926535)
	fmt.Printf("%T\n", err) // *errors.errorString
	// panic(err)
}

func numbersFormatting() {
	fmt.Printf("%e\n", 123456789.123456789) // 1.234568e+08
	fmt.Printf("%g\n", 123456789.12345789)  // 1.234567891234579e+08
	fmt.Printf("%f\n", 123.456)             // 123.456000
	fmt.Printf("%o\n", 25)                  // 31
	fmt.Printf("%x\n", 456789)              // 6x855
	fmt.Printf("%b\n", 42)                  // 101010
	fmt.Printf("%t %t\n", 1 == 0, true)     // false true
}

func stringsFormatting2() {
	fmt.Printf("%s\n", "hello") // hello
	fmt.Printf("%c\n", 'ù')     // ù
	fmt.Printf("%U\n", 'à')     // U+00E0
}
