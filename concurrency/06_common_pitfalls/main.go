package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("===== COMMON PITFALLS & GOTCHAS =====")

	// ============================================================
	// 1. DEADLOCK — Channel (Unbuffered Send with No Receiver)
	// ============================================================
	// A deadlock happens when goroutines are all waiting for each
	// other and NOBODY can make progress.
	//
	// Go detects deadlocks at runtime:
	//   "fatal error: all goroutines are asleep - deadlock!"
	//
	// UNCOMMENT the block below to see the deadlock:
	//
	//   ch := make(chan int)
	//   ch <- 42  // BLOCKS FOREVER! No goroutine to receive.
	//             // Main is the only goroutine — and it's stuck.
	//
	// FIX: Use a goroutine to send or receive:
	//   ch := make(chan int)
	//   go func() { ch <- 42 }()
	//   fmt.Println(<-ch)
	fmt.Println("\n--- Deadlock: Channel (explained in comments) ---")
	fmt.Println("  See comments for deadlock examples (uncomment to try)")

	// ============================================================
	// 2. DEADLOCK — Mutex (Double Lock)
	// ============================================================
	// Locking a mutex that's already locked by the same goroutine
	// causes a deadlock because Go's sync.Mutex is NOT re-entrant.
	//
	// UNCOMMENT to see:
	//   var mu sync.Mutex
	//   mu.Lock()
	//   mu.Lock()  // DEADLOCK! This goroutine already holds the lock.
	//
	// Common in production: Function A locks mu, then calls Function B
	// which also tries to lock mu → DEADLOCK.
	//
	// Fix: Design functions so they don't double-lock, or use
	// internal helper functions that assume the lock is already held.
	fmt.Println("\n--- Deadlock: Mutex (explained in comments) ---")
	fmt.Println("  Go's sync.Mutex is NOT re-entrant (can't lock twice)")

	// ============================================================
	// 3. RACE CONDITION — Without Protection
	// ============================================================
	// Multiple goroutines reading AND writing the same variable
	// without synchronization = RACE CONDITION.
	//
	// Run this with: go run -race concurrency/06_common_pitfalls/main.go
	// to see the race detector catch it!
	fmt.Println("\n--- Race Condition Demo ---")

	// UNSAFE version (has a race condition)
	unsafeCounter := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++ // ❌ RACE! Multiple goroutines read+write without sync
		}()
	}
	wg.Wait()
	fmt.Printf("  Unsafe counter: %d (expected 100, might be wrong!)\n", unsafeCounter)

	// SAFE version (using Mutex)
	safeCounter := 0
	var mu sync.Mutex
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			safeCounter++ // ✅ Protected by mutex
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("  Safe counter:   %d (always 100) ✅\n", safeCounter)

	// SAFE version (using atomic — even faster for simple counters)
	// import "sync/atomic"
	// var atomicCounter int64
	// atomic.AddInt64(&atomicCounter, 1)

	// ============================================================
	// 4. THE RACE DETECTOR
	// ============================================================
	// Go has a built-in race detector. Use it during development!
	//
	//   go run -race main.go
	//   go test -race ./...
	//   go build -race -o myapp
	//
	// It instruments your code and detects concurrent access to
	// shared variables at RUNTIME (not compile time).
	//
	// WARNING: Race detector adds ~10x CPU and memory overhead,
	// don't use in production builds.
	//
	// Interview tip: "Always run tests with -race in CI."
	fmt.Println("\n--- Race Detector ---")
	fmt.Println("  Run: go run -race concurrency/06_common_pitfalls/main.go")
	fmt.Println("  The -race flag detects data races at runtime")

	// ============================================================
	// 5. GOROUTINE LEAK
	// ============================================================
	// A goroutine that's blocked forever and can never exit = LEAKED.
	// It stays in memory forever, consuming resources.
	//
	// Common causes:
	//   - Sending to a channel nobody reads
	//   - Waiting on a channel that's never closed
	//   - Infinite loop with no exit condition
	fmt.Println("\n--- Goroutine Leak ---")

	// BAD: This goroutine leaks! (commented out to not actually leak)
	//   ch := make(chan int)
	//   go func() {
	//       val := <-ch  // Blocks forever — nobody sends to ch!
	//       fmt.Println(val)
	//   }()
	//   // ch is never used again... goroutine is stuck FOREVER.

	// GOOD: Use context to cancel goroutines that are no longer needed.
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("  Goroutine: cleaned up, exiting! ✅")
				return
			default:
				// Do some work...
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(150 * time.Millisecond)
	cancel() // Signal the goroutine to stop
	time.Sleep(100 * time.Millisecond)

	// Interview tip: "Always provide an exit path for goroutines.
	// Use context.Context or a done channel."

	// ============================================================
	// 6. CHANNEL AXIOMS (nil / closed channel behavior)
	// ============================================================
	// These are TRICKY interview questions!
	//
	// ┌───────────┬─────────────────┬─────────────────────────────┐
	// │ Operation │   nil channel   │      closed channel         │
	// ├───────────┼─────────────────┼─────────────────────────────┤
	// │ Send      │ blocks forever  │ PANIC!                      │
	// │ Receive   │ blocks forever  │ returns zero value (+ false)│
	// │ Close     │ PANIC!          │ PANIC!                      │
	// └───────────┴─────────────────┴─────────────────────────────┘
	//
	// Key rules:
	//   - Only the SENDER should close a channel, never the receiver.
	//   - Closing is only needed when the receiver must know no more
	//     values are coming (e.g., to exit a range loop).
	//   - You CAN read from a closed channel (returns zero values).
	//   - You CANNOT send to a closed channel (panics).
	fmt.Println("\n--- Channel Axioms ---")

	// Demo: Reading from a closed channel
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	close(ch)

	val1, ok1 := <-ch
	fmt.Printf("  Read 1: value=%d, ok=%t (has data)\n", val1, ok1)
	val2, ok2 := <-ch
	fmt.Printf("  Read 2: value=%d, ok=%t (has data)\n", val2, ok2)
	val3, ok3 := <-ch
	fmt.Printf("  Read 3: value=%d, ok=%t (closed, zero value!) ⚠️\n", val3, ok3)

	// Demo: The comma-ok pattern to detect closed channels
	// val, ok := <-ch
	// if !ok { fmt.Println("Channel is closed!") }

	// ============================================================
	// SUMMARY — Interview Cheat Sheet
	// ============================================================
	// 1. Use `go run -race` to detect race conditions
	// 2. Go's sync.Mutex is NOT re-entrant (no double-locking)
	// 3. Only the sender should close a channel
	// 4. Always provide an exit path for goroutines (context/done ch)
	// 5. Sending to a closed channel = PANIC
	// 6. Reading from a closed channel = zero value (no panic)
	// 7. nil channel operations block forever
	// 8. Go detects full deadlocks at runtime (all goroutines stuck)
	fmt.Println("\n✅ Review complete! Run with -race flag to catch data races.")
}
