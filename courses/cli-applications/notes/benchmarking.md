# Benchmarking

- Benchmarking is performed natively with Go via the `go test -bench` command
- A benchmark is just a function named `Benchmark*(b *testing.B)` with an argument of `*testing.B`, which is a utility for benchmarks
- To run all benchmarks and skip any test, run this
  ```
  go test -bench . -run ^$
  ```
