# Blocks and shadows

## Shadowing

- **Shadowing** means to declare a variable in an inner scope with the same name as an existing variable from an outer scope
- From the declaration onward, you cannot access the shadowed variable inside the inner scope

## Universe Block

- Built-in types like `int` and `string` are not considered keywords
- Insted they are **predeclared identifiers** in a so-called *universe block* containing the user-defined program
- **DANGER** Built-in type can be shadowed, by clearly they should not be!

## `if` blocks

- `if` and `else if` keywords do not encapsulate the condition in parentheses
- You can declare a variable wich is scoped only in related `if`, `else if` and `else` like this
  ```go
  if n := rand.Intn(10); n == 0 {
      fmt.Println("That's too low")
  } else if n > 5 {
      fmt.Println("That's too big:", n)
  } else {
      fmt.Println("That's a good number:", n)
  }
  // fmt.Println("n:", n) // n undefined here!
  ```
- Please note that `n` is available only inside the conditional blocks, not outside
- *WARNING* Generating random numbers without a seed returns hard-coded values!
