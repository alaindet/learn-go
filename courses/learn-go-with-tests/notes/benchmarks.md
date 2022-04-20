# Benchmarks

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
