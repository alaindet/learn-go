# Flow control

## `if` statement

### Initialization Statement
- An *initialization statement* is a simple statement that can be evaluated by Go before the expression of the if statement
- These simple statements should be used to declare or assign a new variable or to invoke a function
- Variables declared like this are NOT part of the parent scope

Ex.:

```go
	cond := 3

	if a := 1; cond == 1 {
		fmt.Println("Condition 1, a:", a)
	} else if a := 2; cond == 2 {
		fmt.Println("Condition 2, a:", a)
	} else {
    fmt.Println("Condition 3, a: ", a)
  }

  // ERROR: a is not defined!
  // fmt.Pritln(a)

  // Outputs
  // Condition 3, a: 2

```

## `switch` statement

- You can match multiple cases
```go
a := 'a'
switch a {
case 'a', 'A':
    fmt.Println("This is an A")
    break
default:
    fmt.Println("Letter not found")
}
```

- Since Go does not automatically fall through switch cases, you can force that with `fallthrough`
```go
product := "Kayak"
for i, c := range product {
    switch c {
    case 'K':
        fmt.Println(i, "Uppercase character")
        fallthrough // Continue checking cases!
    case 'k':
        fmt.Println(i, "It is a k at #", i)
    default:
        fmt.Println(i, "Letter:", string(c))
    }
}
// 0 Uppercase character
// 0 It is a k at # 0
// 1 Letter: a
// 2 Letter: y
// 3 Letter: a
// 4 It is a k at # 4
```

- With an initialization statement
```go
for counter := 0; counter < 20; counter++ {
    switch val := counter / 2; val {
    case 2, 3, 5, 7:
        fmt.Println("Prime value:", val)
    default:
        fmt.Println("Non-prime value:", val)
    }
}
```

- Removing *comparison value* allows you to use *expressions* instead of just values
```go
for counter := 0; counter < 10; counter++ {
    switch {
    case counter == 0:
        fmt.Println("Zero value")
    case counter < 3:
        fmt.Println(counter, "is < 3")
    case counter >= 3 && counter < 7:
        fmt.Println(counter, "is >= 3 && < 7")
    default:
        fmt.Println(counter, "is >= 7")
    }
}
```
