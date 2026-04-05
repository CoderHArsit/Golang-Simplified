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

### 3. Select
`select` lets a goroutine wait on multiple channel operations. It's like a `switch` statement for channels.

### 4. WaitGroups
`sync.WaitGroup` coordinates goroutines — waits for a collection of goroutines to finish before proceeding.

### 5. Mutexes
`sync.Mutex` protects shared data from race conditions when multiple goroutines read/write the same variable.

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
```

## 🏭 Advanced Patterns

### 6. Worker Pool
A fixed number of goroutines process tasks from a shared jobs channel. Prevents spawning too many goroutines.
```
[Producer] → jobs ch → [Worker 1]
                        [Worker 2] → results ch → [Collector]
                        [Worker 3]
```

### 7. Pipeline
Data flows through stages connected by channels. Each stage transforms data and passes it downstream.
```
[generator] → [square] → [print]
```

### 8. Context
The standard way to cancel goroutines, set deadlines, and manage timeouts.
- `context.WithCancel()`: Manual cancellation signal.
- `context.WithTimeout()`: Auto-cancel after a duration.

### 9. Rate Limiter
Control how frequently events are processed using `time.Tick`.

### 10. Semaphore
A buffered channel used to limit the max number of concurrent goroutines.
```go
sem := make(chan struct{}, 3) // Max 3 concurrent
```
