# Concurrency

- Go was built around the concept of concurrency, which is then native to Go
- A *goroutine* is a thread of execution
- By definition, **concurrency** is the act of loading 2+ more *goroutines* at the same time. If one *goroutine* stops, another one is picked up and started
- Single-core CPU can run only ONE concurrent application
- **parallelism** is the execution of 2+ *goroutines* at the same time, it required multi-core CPUs necessarily

Concurrency, in general, deals with independently execution of processes, while parallelism is just a consequence of it.

## goroutine

- A goroutine is a lightweight thread of execution
- A goroutine is a function, which can run concurrently with other functions
- The keyword `go` is used to declare a goroutine
- They just simulate a thread and **occupy down to 2kb** as compared to 2Mb of normal threads
- A normal thread is of fixed size, while goroutines can resize themselves
- Goroutines are managed by some internal *Go Scheduler* as opposed to the OS native scheduler
- The Go scheduler follows a `m:n scheduling` pattern (TODO), so that m goroutines can be mapped to n OS threads
- There is **no identity** of goroutines, they cannot be accessed or compared
