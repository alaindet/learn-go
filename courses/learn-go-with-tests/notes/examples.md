# Examples

- Test files (`*_test.go`) can contain examples
- An example is a function referring to another function exclusively
- Example function should be `ExampleFoo` for function `Foo`
- An example does not directly test output, but shows how a function is used
- In order for an example to work you need to provide semantic comments
- Valid examples appear in `godoc`, (ex.: `godoc -http=:8080`)

```go
// adder_test.go

// This is run as it has semantic comments
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// This is ignored
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
```
