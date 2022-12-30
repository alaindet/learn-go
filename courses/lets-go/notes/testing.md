# Testing

- Like most things in Go, testing is simple and built-in
- You basically write Go code which controls some other Go code, simple as that
- Some small conventions are needed, but there's no strict necessity of third-party software to test your code, although such libraries exist

## Go Conventions
- Any `*_test.go` file is considered to be containing tests and is excluded from builds
- Tests should be as close as possible to the tested code, with the same file names except for the `*_test.go` suffix
  - Ex.: Code in `application.go` should be tested by code in `application_test.go` in the same folder
- Tests are contained in functions starting with `Test*`, like `TestFriendlyDate`
- Test functions must accept only one argument of type `*testing.T` from the `testing` standard library's package
- Use `testing.T.Errorf()` to signal and log that a test has failed; it works like `fmt.Printf()`
- Recording a failed test via `testing.T.Fatal()` stops execution of current test/subtest, while recording a failure via `testing.T.Errorf()` keeps executing the test
- If you pass the `-fastfail` flag to `go test`, any failure stops execution of tests only for that package
- Some basic test setup could be
  ```go
  // silly_sum.go file
  func SillySum(a, b int) int {
    return a + b
  }

  // silly_sum_test.go file
  import "testing"

  func TestSillySum(t *testing.T) {
    input := []int{10, 20}
    expected := 30
    result := SillySum(input[0], input[1])
    assert := result == expected

    if !assert {
      t.Errorf("Expected %d but got %+v instead", expected, result)
    }
  }
  
  ```

## Theory snippet: Testing Pyramid
- It's a recommendation for the distribution of tests of a software
- Allows for the best coverage and effort
- The opposite is the terrible *ice-cream cone*, where there are many end-to-end tests and fewer unit tests
- Ex.: Here's more of a trapezoid, but you get the point
  ```
           /‾‾‾‾‾‾‾‾‾‾‾‾‾‾\          ^
          /                \         | Slower
         / End-to-end Tests \        |
        /                    \
       /‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾\
      /    Integration Tests   \
     /                          \
    /‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾\   |
   /          Unit Tests          \  | Faster
  /________________________________\ v
  ```
- You should have many small and fast unit tests
- Less integration tests between components
- As few end-to-end tests as possible, just to cover the required use cases

## Theory snippet: Testing Doubles

- Testing requires faking dependencies of the tested code
- Dependencies are replaced by a variety of **testing doubles**
- Testing doubles are broadly split into 3 to 5 categories, which varies in definition among literature, but they all fulfill some tested code depedency
  - **Stub**: A stub is some static code with hard-coded values used as a dependency just to make a test run; returns static values
  - **Spy**: Has the same API as the replaced dependency, but measures ("spies") the behavior of the calling code (ex.: counts the functions calls, checks arguments); they are not very much used
  - **Mock**: It's similar to a stub, but semi-static values can be conditionally calculated; it can process arguments, it can spy tested code like a spy
  - **Fake**: Simplified concrete implementation of the replaced dependency, like a fake AWS S3 implementation; should only be used when a mock or less is not feasible

## Theory snippet: Arrange-Act-Assert (AAA) or Given-When-Then (GWT)

- It's a pattern to write a good test, no matter the language or the type
- GWT is the convention used for **Behavior-Driven Development (BDD)**
- **Arrange/Given** phase: declare inputs, mocks, stubs etc. to setup the test
- **Act/When** phase: executes the actual test by interacting with the tested code, collects results
- **Assert/Then** phase: compares results with expected outcome by making assertions ("this should be like this"), determines the success or failure of the test
