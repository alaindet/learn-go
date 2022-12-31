# Time

- Time is managed via the `time` standard package
- The **Reference Time** for any time formatting in Go is `Monday, 02-Jan-06 15:04:05 MST` just because values are all different and sequential
- From the official documentation (https://pkg.go.dev/time#pkg-constants)
  ```go
  const (
    // The reference time, in numerical order.
    Layout = "01/02 03:04:05PM '06 -0700"
    // ...
  )
  ```
