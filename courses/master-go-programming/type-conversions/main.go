package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.6

	// Compile time error
	// x = x * y // MismatchedTypes error

	x = x * int(y)              // converting with int() truncates digits, does not round
	fmt.Println(x)              // 9
	fmt.Printf("%T %T\n", x, y) // int float64

	x = int(float64(x) * y)
	fmt.Println(x) // 32

	var a = 5
	var b int64 = 2
	fmt.Printf("%T %T\n", a, b) // int int64 (These two types are different)
	// a = b // Cannot assign values with two different types (even if comparable)

	myStr1 := string(90) // Convert number to string (string is a Unicode char)
	fmt.Println(myStr1)  // 'Z'
	// myStr2 := string(42.3) // Error: string() can only convert int32
	var myStr3 = fmt.Sprintf("%f", 44.2)
	fmt.Printf("%T %v\n", myStr3, myStr3) // string 44.200000

	var myChar1 = string(34234)             // string() returns a string of one symbol, not a rune!
	fmt.Printf("%T %v\n", myChar1, myChar1) // string 薺

	myStr4 := "3.123"
	fmt.Printf("%T\n", myStr4) // string

	// strconv exposes Parse* methods to convert from strings to *
	var myFloat1, err = strconv.ParseFloat(myStr4, 64)
	_ = err
	fmt.Println(myFloat1 * 1.2) // 3.7476000000000003

	i, err := strconv.Atoi("123") // ATOI = ASCII to integer
	_ = err
	myStr5 := strconv.Itoa(123) // ITOA = integer to ASCII

	fmt.Printf("i type is %T, value is %v\n", i, i)              // i type is int, value is 123
	fmt.Printf("myStr5 type is %T, value is %q", myStr5, myStr5) // myStr5 type is string, value is "123"
}
