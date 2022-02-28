# Constants

## Untyped
- **untyped constants** limit the inference of the Go compiler allowing for more flexible operatorations between data types
- They are a very useful feature to prototype code and "relax" the compiler (not recommended)
- Literal values are untyped constant not assigned

Ex.:
```go
// With typed constant
const price float32 = 275.00
const quantity uint8 = 3
total := float32(quantity) * price // Conversion is needed!
```

```go
// With untyped constant
const price float32 = 275.00
const quantity  = 3
// This is ok, quantity type is int,
// But it can be used as float32 by Go directly
total := quantity * price
```

*WARNING*: Untyped constants allow for a less strict operations, but *untyped variables* cannot do that

```go
var price = 275.00
var tax float32 = 27.50
// ERROR: invalid operation: price + tax (mismatched types float64 and float32)
fmt.Println(price + tax)
```

## Multiple declarations

```go
const price, tax float32 = 275, 20.3 // Typed
const quantity, inStock = 3, true // Untyped
```

## `iota`

The `iota` keyword can be used to create multiple untyped integer constants

```go
const (
  First = iota // 0
  Second // 1
  Third // 2
)
```
