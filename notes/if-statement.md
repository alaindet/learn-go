# If statement

## Initialization Statement
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
