# Chapter 7: Concurrency ⚡

Go was **built for concurrency**. It's one of the primary reasons the language was created at Google.

## 🗒️ Key Concepts

### 1. Goroutines
A goroutine is a lightweight thread managed by the Go runtime (not the OS).
- Created with the `go` keyword: `go myFunction()`
- Costs ~2KB of stack (vs ~1MB for an OS thread).
- Go can run **millions** of goroutines simultaneously.

### 2. Channels
Channels are Go's way of communicating between goroutines (CSP model: Communicating Sequential Processes).
- **Unbuffered**: Sender blocks until receiver is ready (synchronous handshake).
- **Buffered**: Sender blocks only when the buffer is full.
- **Channel Axioms**: Sending to a closed channel panics. Reading from a closed channel returns zero value. nil channel operations block forever.

### 3. Select
`select` lets a goroutine wait on multiple channel operations. It's like a `switch` statement for channels.

### 4. WaitGroups
`sync.WaitGroup` coordinates goroutines — waits for a collection of goroutines to finish before proceeding.

### 5. Mutexes
`sync.Mutex` protects shared data from race conditions when multiple goroutines read/write the same variable.

### 6. Concurrency vs Parallelism
- **Concurrency** = design (structuring your program to handle multiple things)
- **Parallelism** = execution (physically running on multiple CPU cores)
- `runtime.GOMAXPROCS()` controls how many OS threads run goroutines (default = NumCPU)

## ⚠️ The Golden Rule
> **Don't communicate by sharing memory; share memory by communicating.** — Go Proverb
>
> Use channels to pass data between goroutines instead of using shared variables with locks.

## 🚀 How to Run
```bash
go run concurrency/01_goroutines/main.go
# or
go run concurrency/02_channels/main.go
# or
go run concurrency/03_select_waitgroups/main.go
# or
go run concurrency/04_worker_pool/main.go
# or
go run concurrency/05_patterns/main.go
# or
go run concurrency/06_common_pitfalls/main.go
# or
go run concurrency/07_sync_extras/main.go

# Run with race detector (important for 06!)
go run -race concurrency/06_common_pitfalls/main.go
```

## 🏭 Advanced Patterns

### 7. Worker Pool
A fixed number of goroutines process tasks from a shared jobs channel. Prevents spawning too many goroutines.
```
[Producer] → jobs ch → [Worker 1]
                        [Worker 2] → results ch → [Collector]
                        [Worker 3]
```

### 8. Pipeline
Data flows through stages connected by channels. Each stage transforms data and passes it downstream.
```
[generator] → [square] → [print]
```

### 9. Context
The standard way to cancel goroutines, set deadlines, and manage timeouts.
- `context.WithCancel()`: Manual cancellation signal.
- `context.WithTimeout()`: Auto-cancel after a duration.

### 10. Rate Limiter
Control how frequently events are processed using `time.Tick`.

### 11. Semaphore
A buffered channel used to limit the max number of concurrent goroutines.
```go
sem := make(chan struct{}, 3) // Max 3 concurrent
```

## 🔥 Common Pitfalls (Interview Critical)

### 12. Race Conditions & the Race Detector
- Multiple goroutines accessing shared data without sync = **race condition**
- Detect with `go run -race` or `go test -race`
- Fix with `sync.Mutex`, `sync.RWMutex`, or channels

### 13. Deadlocks
- Unbuffered channel send with no receiver
- Double-locking a non-reentrant `sync.Mutex`
- Go detects full deadlocks at runtime: `fatal error: all goroutines are asleep - deadlock!`

### 14. Goroutine Leaks
- A goroutine blocked forever on a channel = leaked memory
- Always provide an exit path: `context.WithCancel` or a `done` channel

### 15. Channel Axioms
| Operation | nil channel | closed channel |
|-----------|-------------|----------------|
| Send | blocks forever | **panic** |
| Receive | blocks forever | returns zero value |
| Close | **panic** | **panic** |

## 🧰 Sync Extras

### 16. sync.Once
Runs a function exactly once, even from multiple goroutines. Used for singleton/lazy init.

### 17. sync.RWMutex
Multiple concurrent readers OR one exclusive writer. Best for read-heavy workloads.

### 18. sync.Map
Concurrent-safe map. Use when keys are stable and many goroutines access disjoint keys.

### 19. errgroup (`golang.org/x/sync`)
`WaitGroup` + error handling. Returns the first error from any goroutine. Production standard.
```go
g := new(errgroup.Group)
g.Go(func() error { return doWork() })
if err := g.Wait(); err != nil { ... }
```

