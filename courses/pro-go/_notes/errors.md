# Errors

- Errors in Go all implement the `error` built-in interface, defined like this
  ```go
  type error interface {
      Error() string
  }
  ```
- A `panic` function is used when generating a fatal blocking error
- The `errors` standard package has functions to create and manage errors
