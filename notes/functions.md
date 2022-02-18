# Functions

- It's a block of reusable code which can accept inputs (arguments)
- The idiomatic way of writing function names is **camelCase**
- Functions must be unique within the same *package*
- Functions and methods can return **multiple values** in Go
- Functions **main()** and **int()** are *predefined* and are called automatically
- There is **no function overloading** in Go
- There is **no default arguments** in Go
- Any code after a return statement inside a block will generate an error

## Arguments

- Everything is passed by value, including function arguments (source https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
- Each variable defined in a Go program occupies a unique memory location

## Return values

- GO can return **multiple function values** unlike other programming languages
- Return types of all returned values must be declared
  ```go
  func sumAndMultiply(a int, b int) (int, int) {
    return a + b, a * b
  }
  ```
- By convention, if there is an error value, it should be the *last returned value*
- Return values can be **named** if you declare them in the return types
  - Named return values are initialized with zero values based on type
  - You can then change them inside the function
  - When your reach a **naked return** like `return`, named values are returned
  - Please avoid naked returns in any function expect very short ones
  ```go
  func mySum(a, b int) (s int) {
    s = a + b
    return
  }
  s1 := mySum(2, 3)
  fmt.Println(s1) // 5
  ```
