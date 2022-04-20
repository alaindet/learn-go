# Testing

Testing is performed by running this inside the module

```
go test
```

You can also run this to run specific functions and subtests

```
# Run a test function
go test -run TestArea

# Run a subtest inside a test function
go test -run TestArea/Rectangle
```

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

## Workflow
The usual workflow for tests is
1. Write a test according to specifications, it will initially fail
2. Write smallest amount of code to pass test (<-- Debatable!)
3. Run test against code, it could fail
4. Loop on 2. and 3. until test passes
5. Refactor, keep running tests to assure refactoring does not break anything

## Principles
- Test public functions, not private ones
