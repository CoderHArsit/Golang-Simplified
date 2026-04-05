package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("===== ADVANCED CONCURRENCY PATTERNS =====")

	// ============================================================
	// 1. PIPELINE PATTERN
	// ============================================================
	// Data flows through a series of stages connected by channels.
	// Each stage is a goroutine that:
	//   - Receives values from an upstream channel
	//   - Performs a transformation
	//   - Sends results to a downstream channel
	//
	// [generator] → [square] → [print]
	fmt.Println("\n--- Pipeline Pattern ---")

	nums := generator(1, 2, 3, 4, 5)
	squared := square(nums)

	for result := range squared {
		fmt.Printf("Pipeline output: %d\n", result)
	}

	// ============================================================
	// 2. CONTEXT — Cancellation & Timeouts
	// ============================================================
	// context.Context is the standard way to:
	//   - Cancel goroutines from the outside
	//   - Set deadlines and timeouts
	//   - Pass request-scoped values
	fmt.Println("\n--- Context Cancellation ---")

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("  Worker: received cancel signal, stopping.")
				return
			default:
				fmt.Printf("  Worker: tick %d\n", i)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(350 * time.Millisecond)
	cancel() // Signal the goroutine to stop
	time.Sleep(50 * time.Millisecond)

	// ============================================================
	// 3. CONTEXT WITH TIMEOUT
	// ============================================================
	fmt.Println("\n--- Context with Timeout ---")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel2() // Always call cancel to release resources

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Work completed")
	case <-ctx2.Done():
		fmt.Println("Timeout! Context expired:", ctx2.Err())
	}

	// ============================================================
	// 4. RATE LIMITER
	// ============================================================
	// Control how frequently events are processed using time.Tick.
	fmt.Println("\n--- Rate Limiter ---")

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(100 * time.Millisecond) // 1 request per 100ms

	for req := range requests {
		<-limiter // Wait for the next "tick"
		fmt.Printf("Request %d processed at %s\n", req, time.Now().Format("15:04:05.000"))
	}

	// ============================================================
	// 5. SEMAPHORE (Limiting Concurrency)
	// ============================================================
	// A buffered channel can act as a semaphore to limit
	// the number of goroutines running simultaneously.
	fmt.Println("\n--- Semaphore (Max 2 concurrent) ---")

	sem := make(chan struct{}, 2) // Max 2 goroutines at a time
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}         // Acquire slot (blocks if 2 are running)
			defer func() { <-sem }()  // Release slot when done

			fmt.Printf("  Task %d: running\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("  Task %d: done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All semaphore tasks completed!")
}

// ============================================================
// Pipeline Helper Functions
// ============================================================

// generator produces values and sends them into a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square reads from 'in', squares each value, and sends to 'out'
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
