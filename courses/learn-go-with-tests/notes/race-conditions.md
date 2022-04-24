# Race condition

A race condition is a situation where the order of operations is not predictable and two goroutines try to read from or write into the same memory address at the same time. Go doc says it's an [unsynchronized accesses to shared variables](https://go.dev/blog/race-detector)

To try to detect race conditions, use the race detector built-in with Go with

```
go test -race
```

Other ways to use the built-in race detector

```
go test -race mypkg    // test the package
go run -race mysrc.go  // compile and run the program
go build -race mycmd   // build the command
go install -race mypkg // install the package
```
