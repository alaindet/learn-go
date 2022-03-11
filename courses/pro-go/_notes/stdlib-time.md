# Time
Most operations with time and dates are managed via the `time` standard package

## Formatting
Formatting time in Go is done either via standard layouts or via a custom string layout which must reference a **specific point in time** which has convenients numbers

`2006/01/02 15:04:05 -0700`

The reason is that rearranging the values you get

`01/02 03:04:05PM '06 -0700`

which inherently has the numbers `1`, `2`, `3`, `4`, `5`, `6` and `7` in sequence, easy to parse for Go

Example

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  layout := "Day is: 02, Month is: 01, Year is: 2006"
  fmt.Println(t.Format(layout)) // Day is: 11, Month is: 03, Year is: 2022
}
```

## Duration
- The `time.Duration` type is an alias of `int64` used to represent time durations in milliseconds
- Durations are widely used by the `time` package to change dates by adding/subtracting time
