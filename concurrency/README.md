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

---

## 🎯 Interview Questions

### Q1: What is a goroutine? How is it different from an OS thread?
**Answer:** A goroutine is a **lightweight thread** managed by the Go runtime, not the OS. Key differences:
- **Stack size**: Goroutine starts at ~2KB (grows dynamically) vs ~1MB for an OS thread.
- **Scheduling**: Go uses an M:N scheduler (many goroutines on few OS threads) — cheaper context switches.
- **Creation**: `go func()` is trivial; millions of goroutines are practical.

### Q2: What is the difference between concurrency and parallelism?
**Answer:**
- **Concurrency** = program **design** — structuring code to handle multiple tasks (may run on 1 core).
- **Parallelism** = **execution** — physically running tasks simultaneously on multiple CPU cores.
- Go makes concurrency easy with goroutines; parallelism depends on `runtime.GOMAXPROCS()` and available cores.

### Q3: What are channels? What's the difference between buffered and unbuffered?
**Answer:**
- **Unbuffered** (`make(chan int)`) — sender blocks until a receiver is ready. It's a synchronous handshake.
- **Buffered** (`make(chan int, 5)`) — sender blocks only when the buffer is full. Allows async communication up to the buffer size.
- Channels implement Go's CSP model: "Don't communicate by sharing memory; share memory by communicating."

### Q4: What happens when you send to a closed channel? Read from a closed channel? Close a nil channel?
**Answer:**
| Operation | nil channel | closed channel |
|-----------|-------------|----------------|
| **Send** | blocks forever | **PANIC** |
| **Receive** | blocks forever | returns zero value + `false` |
| **Close** | **PANIC** | **PANIC** |

Rule: Only the **sender** should close a channel, never the receiver.

### Q5: What is a deadlock? How does Go detect it?
**Answer:** A deadlock occurs when all goroutines are blocked and no progress can be made. Go detects **full deadlocks** at runtime with: `fatal error: all goroutines are asleep - deadlock!`. Common causes: unbuffered send with no receiver, double-locking a `sync.Mutex` (not re-entrant), circular channel dependencies.

### Q6: What is a race condition and how do you detect it?
**Answer:** A race condition occurs when multiple goroutines access shared data concurrently and at least one writes, without synchronization. Detect with Go's built-in **race detector**: `go run -race main.go` or `go test -race ./...`. Fix with `sync.Mutex`, `sync.RWMutex`, channels, or `sync/atomic`.

### Q7: How does `select` work? What happens with a `default` case?
**Answer:** `select` waits on multiple channel operations and executes the first one that's ready. If multiple are ready, one is chosen **randomly**. Adding a `default` case makes it **non-blocking** — if no channel is ready, `default` executes immediately. Without `default`, `select` blocks until a case is ready.

### Q8: What is `context.Context` and why is it important?
**Answer:** `context.Context` is the standard mechanism for:
- **Cancellation**: `context.WithCancel()` — propagate cancel signals to goroutines.
- **Timeouts**: `context.WithTimeout()` — auto-cancel after a duration.
- **Deadlines**: `context.WithDeadline()` — cancel at a specific time.
- **Values**: `context.WithValue()` — pass request-scoped data.
It should be the **first parameter** of any function that does I/O or long-running work.

### Q9: What is a goroutine leak and how do you prevent it?
**Answer:** A goroutine leak happens when a goroutine is blocked forever (on a channel or lock) and can never exit. It stays in memory permanently. Prevention:
- Always provide an **exit path** (context cancellation, done channel).
- Always **close channels** when done sending.
- Use `context.WithTimeout` for operations that might hang.

### Q10: What is the Worker Pool pattern and why is it useful?
**Answer:** A fixed number of goroutines (workers) read tasks from a shared jobs channel, preventing unbounded goroutine creation. Structure:
```
[Producer] → jobs channel → [Worker 1, 2, 3] → results channel → [Collector]
```
Benefits: controlled resource usage, backpressure, and predictable memory footprint. In production, use `errgroup.SetLimit()` for a simpler version.

### Q11: What is the difference between `sync.Mutex` and `sync.RWMutex`?
**Answer:**
- `sync.Mutex` — exclusive lock; only **one** goroutine (reader or writer) at a time.
- `sync.RWMutex` — allows **multiple concurrent readers** OR **one exclusive writer**.
Use `RWMutex` for read-heavy workloads (caches, configs). Use `Mutex` when reads ≈ writes.

### Q12: What is `sync.Once` and what's a common use case?
**Answer:** `sync.Once` ensures a function runs **exactly once**, even when called from multiple goroutines concurrently. Common use case: **singleton pattern** (lazy initialization of database connections, loggers, or configuration). It's goroutine-safe and more idiomatic than manual locking.
