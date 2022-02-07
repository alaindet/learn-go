package main

import "fmt"

/*
Constants are checked at compile-time, so any error related to constants
is detected before running the program

- They must be in camelCase
- They must be initialized when declared
- They are checked at compile-time
- They throw a warning and not an error if not used

Rules
1. Constants cannot be changed
   ```go
   const myConst = 5
   myConst = 6 // Error
   ```
2. Constants cannot be initialized with a function call
   ```go
   const myConst = math.Pow(2, 3) // Error
   ```
3. Constants cannot be initialized with a variable
   ```go
   myVar := 5
   const myConst = myVar
   ```
4. Constants can be initialized with the len() function on a string literal
   ```go
   const myConst = len('this is a string literal')
   ```
*/
func main() {
	// const daysInWeek int = 7
	// const pi float64 = 3.14159
	const secondsInHour int = 60 * 60

	durationInHours := 24
	fmt.Printf("Duration in seconds: %d\n", durationInHours*secondsInHour)

	// Run-time error example
	// x1, y1 := 5, 0
	// fmt.Println(x1 / y1) // This error is checked only at run-time

	// Compile-time error example
	// const x2 = 5
	// const y2 = 0
	// fmt.Println(x2 / y2) // Compiler throws a DivByZero error

	// Multiple declarations
	const a1, b1 int = 4, 5 // Explicit type
	const a2, b2 = 6, 7     // Implicit type
	const (
		a3 = 8
		b3 = 9
		c3 = 10
	)

	// In a grouped declaration, variables without type and initial value "borrow"
	// the last known type and initial value in the list, meaning that
	// c4 is initialized the same as b4
	//
	// Example
	//
	// const (
	// 	ax = 42
	// 	bx
	// 	cx
	// )
	//
	// ax, bx and cx are all initialized as ax (int 42)
	//
	const (
		a4 = 42
		b4 = 43
		c4
	)
	fmt.Println(a4, b4, c4) // Prints 42, 43, 43

	const a float64 = 5.2      // Typed constant
	const b = 6.7              // Untyped constant
	const c float64 = a * b    // Types constant
	const d = "Hello " + "Go!" // Untyped constant (inferred from expression)
	const e = 5 > 6            // Untyped constant (inferred from expression)
	const f = 2                // Untyped constant (inferred from expression)
	const g = a * f            // Untyped constant (cast to float64)
	fmt.Println(a, b, c, d, e)
	fmt.Printf("g constant type is: %T\n", g)

	var myVar1 int = f     // x constant is untyped, gets casted to int
	var myVar2 float64 = f // x constant is untyped, gets casted to float64
	fmt.Printf("myVar1 type is: %T, myVar2 type is %T\n", myVar1, myVar2)
}
