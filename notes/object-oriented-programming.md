# Object Oriented Programming (OOP)
- In Go there is no classes, but you can define **methods** on *structs*

## Methods
- A **receiver** is a type on which you can declare a function with special context
- A **method** or *receiver function* is like a class method
- A method takes the receiver (the type) *as an argument*
- Syntax is like this
  ```go
  type names []string

  func (n names) print() {
    for i, name := range n {
      fmt.Println(i, name)
    }
  }
  ```
- As a convention, avoid using `this` or `self` as the name of the argument containing the type
- Unless confusing, try to use the first letter of the type, ex.: `n` for type `names`
- It is possible to call

## Methods and functions
- A function is scoped inside a package
- A method is scoped inside a type
- A method takes the receiver type as its first argument and follows a different syntax
- A method can be called from the type itself or from a variable having that type
- Method declarations cannot be created if the receiver is a named pointer type
  ```go
  type distance *int

  // Error: Invalid receiver type
  func (d *distance) getMessage() {
    return "Here is a message""
  }
  ```

## Conversions
- It's idiomatic to convert the type of an expression to convert it via a specific method
- Example
  ```go
  import (
    "time"
    "fmt"
  )

  // ...
  var n int64 = 93422433
  fmt.Println(n)
  fmt.Println(time.Duration(n))
  ```

## Conventions
- If you're changing the receiver data and want to save memory, use a pointer as receiver type
- Otherwise, use a normal value for receiver
  ```go
  // Normal value receiver
  func (c car) changeCar(newBrand string, newPrice float64) string {
    return c.brand
  }

  // Pointer value receiver
  func (c *car) changeCar(newBrand string, newPrice float64) {
    (*c).brand = newBrand
    (*c).price = newPrice
  }
  ```
