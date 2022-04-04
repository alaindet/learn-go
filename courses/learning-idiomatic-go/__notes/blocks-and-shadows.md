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
- Iterates on every iterable built-in type, like *arrays*, *slices*, *strings*, *maps* and *channels*, as well as compound types derived from said types
- Returns two variables per iteration: key and value for maps, index and value for everything else
- Assigned values are **COPIED** not referenced
- Order of keys is usually random for security reason (*Hash DoS attacks*)
- Order is still preserved alphabetically when printing the whole map data

```go
// Slices
for index, value := range []int{1, 3, 5, 7, 9} {
    fmt.Println(index, value)
}

// Maps
for key, value := range map[string]int{ "Alice": 20, "Bob": 30, "Charlie": 40 } {
    fmt.Println(key, value)
}

// Strings
// NOTE: letter is a rune, so it's an int32
// range on strings loops over runes and groups multi bytes as neeeded
word := "the_π_is_cool"
fmt.Printf("len(%q): %d\n", word, len(word))
for index, letter := range word {
    fmt.Println(index, letter, string(letter))
}
// len("the_π_is_cool"): 14 // <-- NOTE: 14 is bytes length, 13 is the runes count!
// 0 116 t
// 1 104 h
// 2 101 e
// 3 95 _
// 4 960 π <-- NOTE: This is 2-bytes long (>128), index 5 is merged here and SKIPPED!
// 6 95 _
// 7 105 i
// 8 115 s
// 9 95 _
// 10 99 c
// 11 111 o
// 12 111 o
// 13 108 l

// Ignoring index
for _, value := range []int{1, 3, 5, 7, 9} {
    fmt.Println(index, value)
}

// Ignoring value
for key := range map[string]int{ "Alice": 20, "Bob": 30, "Charlie": 40 } {
    fmt.Println(key)
}
```

## `switch` statements
- `switch` can be used on *expressions* as well as on **types*
- `switch` accepts simple statements like `if` does
- `case` can have multiple conditions, but there is no implicit fallthrough
- Cases create a block and stop when matched, no need for `break`
- If you really want to use fallthrough mechanism, you can use the `fallthrough` keyword at the end of a case to make it fall through the next case
- Explicit `fallthrough` is rare and should be avoided

```go
words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

for _, word := range words {
    switch size := len(word); size {
    case 1, 2, 3, 4:
        fmt.Println(word, "is a short word!")
    case 5:
        wordLen := len(word)
        fmt.Println(word, "is exactly the right length:", wordLen)
    case 6, 7, 8, 9:
    default:
        fmt.Println(word, "is a long word!")
    }
}
// a is a short word!
// cow is a short word!
// smile is exactly the right length: 5
// anthropologist is a long word!
```
