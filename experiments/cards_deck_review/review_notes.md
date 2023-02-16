# Review notes

01. Snake case is not idiomatic of Go
02. Deck methods should be public (capital letter)
03. Naked returns are discouraged for errors, slices and structs since you don't have control over initialization and can also lead to complex code
04. `main()` should be the first function in the code since any function in a package is hoisted in Go: i.e. functions definition bubble up so you can use functions you have not declared yet
05. Add a `NewDeck` constructor function instead of the `create` method on an existing instance of `Deck`
06. Add types for suits and values
07. Avoid using `int64` for indexing, convert `int64` to `int` instead
08. `suits` and `values` should be package-level variables (they cannot be constants yet as of Go 1.20)
09. `map[int64][string]` is just equivalent to `[]string` if using `int` as key
10. There is no performance gain in saving `len(something)` in a variable, since the length of variables is already stored as readonly in the stack
11. Prefer a separate function to extract random numbers in a range, like the `getRandomInt` provided function
12. Avoid removing cards from original deck when shuffling; this way you can shuffle multiple times
13. `make(type, len, cap)` should specify the `cap` when creating slices of known size to avoid reallocating memory when `append`ing
14. Create known errors as global variables, like `ErrInvalidQuantity` and `ErrInsufficientCards`
15. `check` method does not convey meaning, while `CheckAvailableCards` is more explicit
16. Checking for `err == nil` instead of `err != nil` leads to unwanted indentation for the happy path, instead of catching errors and exit early
17. `error` implements the `Stringer` interface, so `err.Error()` is not necessary when printing, i.e.: `fmt.Println(err)` is ok, is `fmt.Println(err.Error())` is redundant
