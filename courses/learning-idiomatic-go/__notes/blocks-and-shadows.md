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

## `for` blocks

- The `for` keyword is the only looping keyword in Go
- However, it can be used in 4 ways
  - Complete `for` loop
  - Condition-only `for` loop
  - Infinite `for` loop
  - `for-range` loop

### Complete `for` loop
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### Condition-only `for` loop
- Equivalent to `while` in other programming languages

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

```

### Infinite `for` loop
- Best for games and listening to channels
- Can only be exited via `break`

```go
for {
    // ...
    if someCondition {
        break
    }
}
```

### `for-range` loop
- The most idiomatic loop in Go, preferred if possible
- Iterates on every iterable built-in type, like arrays, slices and maps, as well as compound types derived from said types
- Returns two variables per iteration: index and value for slices and arrays, key and value for maps

```go
// Slices
for index, value := range []int{1, 3, 5, 7, 9} {
    fmt.Println(index, value)
}

// Maps
for key, value := range map[string]int{ "Alice": 20, "Bob": 30, "Charlie": 40 } {
    fmt.Println(key, value)
}

// Ignoring index
for _, value := range []int{1, 3, 5, 7, 9} {
    fmt.Println(index, value)
}

// Ignoring value
for key := range map[string]int{ "Alice": 20, "Bob": 30, "Charlie": 40 } {
    fmt.Println(key)
}
```
