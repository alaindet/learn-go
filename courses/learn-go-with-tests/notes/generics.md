# Generics

Just some notes on generics, introduced on Go 1.18.

These two functions are not equivalent

```go
func GenericFoo[T any](x, y T)
func InterfaceyFoo(x, y interface{})
```

Even if `any` is an alias for `interface{}`

- With `InterfaceyFoo` you can pass `x` and `y` of different types
- With `GenericFoo` you are forced to pass `x` and `y` of the same type, although it can be any type

- Generics are called *type parameters*
- Generics allows for `comparable` type parameter, which implies any type can be compared natively via `==` and `!=`
- `comparable` is described in the Go Programming Language Specification available at (https://go.dev/ref/spec)[https://go.dev/ref/spec], at the section *Comparison operators*
