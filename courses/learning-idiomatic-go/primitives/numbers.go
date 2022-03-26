package main

import "fmt"

func numberLiteralsExample() {
	fmt.Println("Literals")
	fmt.Println("Type | Expression | Decimal value")
	fmt.Println("Integer | 123 |", 123)
	fmt.Println("Hexadecimal | 0xcafe |", 0xcafe)
	fmt.Println("Octal | 0o123 |", 0o123)
	fmt.Println("Binary | 0b1001 |", 0b1001)
	fmt.Println("With separator | 42_000_000 |", 42_000_000)
	fmt.Println("Floating | 12.3 |", 12.3)
	fmt.Println("Floating (no decimals) | 42. |", 42.)
}
