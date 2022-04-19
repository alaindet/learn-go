# Testing

## Conventions
- Tests need to be in a file named like `*_test.go` (ex.: `hello_test.go`)
- Test functions must start with `Test` (ex.: `TestHello`)
- You must `import "testing"` in the test file
- The test function must accept only one argument `t *testing.T`
- Test files can contain examples (See **Examples** section)

## Basics
In Go, testing is built into the language. In general, `go test` is used to perform tests. These things are assumed

```go
// hello.go
package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

// hello_test.go
package main

import "testing"

func TestHello(t *testing.T) {

	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

## Examples
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

## Benchmarks
- Benchmarks are similar to examples
- Naming is `BenchmarkFoo` for function `Foo`
- Benchmarks are run with `go test -bench=.` or `go test -bench=. -benchmem`
- It accepts a single argument `*testing.B`

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
```

## Workflow
The usual workflow for tests is
1. Write a test according to specifications, it will initially fail
2. Write smallest amount of code to pass test (<-- Debatable!)
3. Run test against code, it could fail
4. Loop on 2. and 3. until test passes
5. Refactor, keep running tests to assure refactoring does not break anything

## Principles
- Test public functions, not private ones
