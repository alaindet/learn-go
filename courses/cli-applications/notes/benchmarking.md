# Benchmarking

- Benchmarking is performed natively with Go via the `go test -bench` command
- A benchmark is just a function named `Benchmark*(b *testing.B)` with an argument of `*testing.B`, which is a utility for benchmarks

- To run all benchmarks and skip any test, run this
  ```
  go test -bench . -run ^$
  ```
- When benchmarks take long (> 1s), Go can decide to run it only once. To force into running it multiple times, run this
  ```
  go test -bench . -run ^$ -benchtime=10x
  ```

- With the CPU profiler enabled
  ```
  go test -bench . -run ^$ -benchtime=10x -cpuprofile cpu00.pprof
  ```

- Analyze the `cpu00.pprof` generated file via
  ```
  go tool pprof cpu00.pprof
  ```

- With the memory profiler enabled
  ```
  go test -bench . -benchtime=10x -run ^$ -memprofile mem00.pprof
  ```

- Analyze the `mem00.pprof` file like this
  ```
  go tool pprof -alloc_space mem00.pprof
  ```

- Create a comparison file for memory usage
  ```
  go test -bench . -benchtime=10x -run ^$ -benchmem | tee benchresults00m.txt
  ```

- Create a comparison file for CPU usage
  ```
  go test -bench . -benchtime=10x -run ^$ | tee benchresults00.txt
  ```

## Comparing benchmarks

- Install `benchstat`
  ```
  go get -u -v golang.org/x/perf/cmd/benchstat
  ```
  From docs: "Benchstat computes statistical summaries and A/B comparisons of Go benchmarks."
  ```
  benchstat results1.txt results2.txt
  benchstat ./profiling/thinkpad/v0/benchresults00m.txt ./profiling/thinkpad/v1/benchresults01m.txt

  benchstat benchresults00m.txt benchresults01m.txt
  ```
