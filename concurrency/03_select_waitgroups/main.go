package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("===== SELECT & WAITGROUPS =====")

	// 1. WaitGroup — The Proper Way to Wait for Goroutines
	// Instead of time.Sleep, use sync.WaitGroup.
	// - wg.Add(n): Tell the group to wait for n goroutines.
	// - wg.Done(): Called by each goroutine when it finishes. (Usually deferred)
	// - wg.Wait(): Blocks until all goroutines call Done().
	fmt.Println("\n--- WaitGroup ---")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add 1 to the counter
		go func(id int) {
			defer wg.Done() // Decrement counter when goroutine finishes
			fmt.Printf("  Worker %d: started\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("  Worker %d: finished\n", id)
		}(i)
	}

	wg.Wait() // Block until all 3 workers call Done()
	fmt.Println("All workers completed!")

	// 2. Select — Waiting on Multiple Channels
	// select picks the FIRST channel that's ready.
	// If multiple are ready, it picks one at RANDOM.
	fmt.Println("\n--- Select ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "result from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "result from ch2"
	}()

	// Wait for the first result
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	}

	// 3. Select with Timeout
	// Prevent waiting forever using time.After.
	fmt.Println("\n--- Select with Timeout ---")
	slowCh := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Takes too long!
		slowCh <- "slow result"
	}()

	select {
	case msg := <-slowCh:
		fmt.Println("Got:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout! Gave up waiting after 500ms.")
	}

	// 4. Mutex — Protecting Shared Data
	// Without a mutex, concurrent writes to the same variable cause a RACE CONDITION.
	fmt.Println("\n--- Mutex ---")
	var mu sync.Mutex
	counter := 0
	var wg2 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()   // Lock: only one goroutine can access counter at a time
			counter++
			mu.Unlock() // Unlock: let others in
		}()
	}

	wg2.Wait()
	fmt.Printf("Counter (with mutex): %d (should be 1000)\n", counter)

	// 5. Fan-Out / Fan-In Pattern
	// Multiple goroutines reading from one channel (fan-out)
	// Multiple results collected into one channel (fan-in)
	fmt.Println("\n--- Fan-Out, Fan-In ---")
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg3 sync.WaitGroup

	// Launch 3 workers (fan-out)
	for w := 1; w <= 3; w++ {
		wg3.Add(1)
		go worker(w, jobs, results, &wg3)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers and close results
	go func() {
		wg3.Wait()
		close(results)
	}()

	// Collect results (fan-in)
	for r := range results {
		fmt.Printf("  Result: %d\n", r)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("  Worker %d processing job %d\n", id, j)
		results <- j * 2
	}
}
