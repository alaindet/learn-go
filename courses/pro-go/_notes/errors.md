# Errors

- Errors in Go all implement the `error` built-in interface, defined like this
  ```go
  type error interface {
      Error() string
  }
  ```
- The `errors` standard package has functions to create and manage errors

## Panicking
- A `panic()` function is used when generating a fatal blocking error
- When a panic error is generated, it triggers all the `defer` statements in the current function, then bubbles up the stack interrupting the execution
- Whether a function should trigger a normal error or a panic error should be decided by the calling function which has context
- Conventionally, such functions should have two version, one of which panics and is prefix with `Must` (Ex.: the `regexp` standard package has `Compile` and `MustCompile` functions)
- The built-in `recover()` function must be called via `defer` and can capture a panic error
  ```go

  ```
