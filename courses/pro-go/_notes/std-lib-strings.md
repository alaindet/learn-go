# Strings
- Most operations on strings can be performed using the standard library packages `strings` and `unicode`
- The `strings.Builder` is more efficient than concatenating when building strings, especially when using the `strings.Builder.Grow` method to allocate memory in advance

## Formatting
Formatting is performed mostly via the `fmt` standard package. The main functions to format output strings are `fmt.Printf()` and `fmt.Sprintf()`. They work to same, but `Printf` prints to standard output, while `Sprintf` returns a string

Formatting is done via a **template** containing some **verbs**, combined with real values

Example
```go
fmt.Printf(
  "Value: %s", // The template (contains %s which is a verb)
  "hello", // The argument(s)
) // Value: hello
```

All verbs are listed here [https://pkg.go.dev/fmt#hdr-Printing](https://pkg.go.dev/fmt#hdr-Printing), common ones are
- `%s`: string
- `%q`: quoted string
- `%d`: digit (an integer number)
- `%f`: floating (a decimal number)
- `%v`: Value in readable format
- `%#v`: Value in Go-syntax format
- `%T`: Type in Go-syntax format

Verbs accept further **specifiers**, for example the `%f` accepts specifiers for how many digits to show before and after the dot

```go
fmt.Printf("Value: %.2f", 3.1415926535) // Value: 3.14
```

- Methods `String` and `GoString` allow developers to change the representation of a struct when using the verbs `%v` and `%#v` respectively
- The built-in corresponding interfaces are called `Stringer` and `GoStringer`
- Note that `String` and `GoString` must be defined in the same package as their type, so native types must be aliased in order to change their `String` and `GoString`

Example
```go
type Person struct {
	Name string
}

// Changes default representation of %v verb for Person
func (p *Person) String() string {
	return fmt.Sprintf("V: Person name: %s", p.Name)
}

// Changes default representation of %#v verb for Person
func (p *Person) GoString() string {
	return fmt.Sprintf("#V: Person name: %s", p.Name)
}

func main() {
	p := &Person{"John"}
	fmt.Printf("%v\n", p) // V: Person name: John
	fmt.Printf("%#v\n", p) // #V: Person name: John
}

```

## Scanning
Scanning means capturing strings from user input or other readers, like a file. Functions can be grouped in 3 groups

- Read from standard input
  - `fmt.Scan(...vals)`: Read everything, newlines are like spaces, split by spaces
  - `fmt.Scanf(template, ...vals)`: Like `Scan()`, but scans through a template
  - `fmt.Scanln(...vals)`: Like `Scan()`, but newline terminates scanning

- Read from a reader (can be a file)
  - `fmt.Fscan(reader, ...vals)`: Equivalent to `Scan()` but uses a reader
  - `fmt.Fscanf(reader, template, ...vals)`: Equivalent to `Scanf()` but uses a reader
  - `fmt.Fscanln(reader, ...vals)`: Equivalent to `Scanln()` but uses a reader

- Read from a value
  - `fmt.Sscan(str, ...vals)`: Equivalent to `Scan()` but uses a direct value
  - `fmt.Sscanf(str, template, ...vals)`: Equivalent to `Scanf()` but uses a direct value
  - `fmt.Sscanln(str, template, ...vals)`: Equivalent to `Scanln()` but uses a direct value
