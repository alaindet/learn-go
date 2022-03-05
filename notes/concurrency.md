# Concurrency

- Go was built around the concept of concurrency, which is then native to Go
- A *goroutine* is a thread of execution
- By definition, **concurrency** is the act of loading 2+ more *goroutines* at the same time. If one *goroutine* stops, another one is picked up and started
- Single-core CPU can run only ONE concurrent application
- **parallelism** is the execution of 2+ *goroutines* at the same time, it requires multi-core CPUs necessarily

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

## WaitGroups

A WaitGroup is a grouping mechanism that waits for **all** goroutines inside the group to finish before moving on

```go
package main

import (
  "fmt"
  "sync"
  "time"
)

func f1(wg *sync.WaitGroup) {
  fmt.Println("f1() started")
  for i := 0; i < 3; i++ {
    fmt.Println("f1(), i = ", i)
    time.Sleep(time.Second)
  }
  fmt.Println("f1() stopped")
  wg.Done() // Signal to WaitGroup that f1() has finished
}

func main() {
  fmt.Println("main() started")
  var wg sync.WaitGroup
  wg.Add(1)
  go f1(&wg)
  wg.Wait()
  fmt.Println("main() stopped")
}

// main() started
// f1() started
// f1(), i =  0
// f1(), i =  1
// f1(), i =  2
// f1() stopped
// main() stopped
```

## Race condition and thread safety
- If two or more goroutines try to read and/or write on the same memory location, a **data race** can happen, meaning the final value of all the operations depends on the unpredictable order of data accessing and writing
- **thread safety** implies mechanisms to avoid any *race condition* and be able to predict the final result
- Go offers a solution called *race detector*, which is enabled with a flag in the CLI and creates a report of detected race conditions, ex.:
  ```
  go run -race main.go
  ```
- The race detector makes "guesses" based on timestamps and code behavior while measuring the run time of the application, which means the run time conditions and realistic workloads are important to catch race conditions
- Solutions to data race include using **mutexes** and **channels**

## Mutex
- Mutex (Mutual Exclusion Object) is an explicit synchronization mechanism for goroutines
- It's a *synchronization primitive*
- Strictly speaking, a mutex is a **locking mechanism** used to synchronize access to a resource. Only one thread can acquire the mutex. It means there is ownership associated with a mutex, and only the owner can release the lock (mutex)
- Mutex can produce deadlocks (anomalous states where threads lock each other indefinetely) and starvation (impossibility of gaining lock due to deadlocks)

## Channels
- It's a synchronization mechanism to communicate between goroutines
- A channel is shared between goroutines and holds only one data type
- In Go, they act similarly to pointers
- Main operations are
  - **send** sends a value through the channel to a goroutine using the corresponding receive command
  - **receive** receives a value through the channel that was previously sent
  - **close** closes the channel (no more sent/received messages), any subsequent *receive* operation will yield the zero value of the channel
- *Unidirectional channels* can only receive or can only send data
  ```go
  c1 := make(<-chan string) // Receive-only
  c2 := make(chan<- string) // Send-only
  ```
- Bidirectional channel can be *restricted* automatically when assigned to functions expecting unidirectional channel, which means that a bidirectional channel can be assigned as a send-only channel to a function

Example
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	c := make(chan string)

	go func(c chan string) {
		fmt.Println("goroutine start")
		time.Sleep(time.Second * 2)
		c <- "Hello World"
    time.Sleep(time.Second * 1)

    // This is never executed, because after you write into the channel, the main
    // function resumes and it finishes before this goroutine reaches here
    // NOTE: The main function takes it all and stops execution
		fmt.Println("goroutine end")
	}(c)

	message := <-c

	fmt.Println("message:", message)
	fmt.Println("main end")
}
// main start
// goroutine start
// [T+2000ms]
// message: Hello World
// main end
```

### Unbuffered channel
```go
ch1 := make(chan int) // Unbuffered channel
ch2 := make(chan int, 3) // Buffered channel
```
- Unbuffered channels are also called *synchronous channels* as any write or read operation **blocks the execution**
- Unbuffered channels can be seen as buffered channels with capacity = 1

### Buffered channel
- The *capacity* of a channel is the number of values it can store before becoming full
- The *sender* blocks only when the buffer is full
- So, the *sender* keeps writing until the buffer is full
- The *receiver* blocks only when the buffer is empty
- So, the *receiver* keeps reading until the buffer is empty

### The `range` keyword
When using `range` on a channel, it subscribes to the channel, reads all available data and remains listening for new values until the channel is closed via `close`. If the channel is not closed, a **deadlock** error arises and the program stops

Example
```go
func sendValue(c chan string, repetitions int, _ int) {
	for i := 0; i < repetitions; i++ {
		c <- "Hello World"
	}
  // Without this, it breaks
	// close(c)
}

func main() {
	c := make(chan string, 10)
	go sendValue(c, 3, 2)
	for val := range c {
		fmt.Println("Value:", val)
	}
}
```

### Select statement
- Groups operations on channels
- A `select` statement has a similar structure to a switch statement, except that the case statements are channel operations
- Runs all channel operations in `case` statements and if none is matched the `default` case is performed
- This allows for non-blocking reads and writes from and into the channel
- `select` statements are executed only once, so they're usually used within loops
- It is used also to access multiple channels at a time
- **WARNING**: since reading from a closed channel gives the zero-value of the channel, closed channel **should be assigned the `nil` value** so that reading cannot be performed and its select case fails
- If no case matches and there is no default statement, the `select` statements **blocks** the execution until some case matches
