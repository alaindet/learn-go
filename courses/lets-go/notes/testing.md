# Testing

- Like most things in Go, testing is simple and built-in
- You basically write Go code which controls some other Go code, simple as that
- Some small conventions are needed, but there's no strict necessity of third-party software to test your code, although such libraries exist

## Conventions
- Any `*_test.go` file is considered to be containing tests and is excluded from builds
- Tests should be as close as possible to the tested code, with the same file names except for the `*_test.go` suffix
  - Ex.: Code in `application.go` should be tested by code in `application_test.go` in the same folder
- Tests are contained in functions starting with `Test*`, like `TestFriendlyDate`
- Test functions must accept only one argument of type `*testing.T` from the `testing` standard library's package
- Use `testing.T.Errorf()` to signal and log that a test has failed; it works like `fmt.Printf()`
- Failed test **do not** halt execution of the remaining tests
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
