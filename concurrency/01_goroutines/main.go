package main

import (
	"fmt"
	"time"
)

// A regular function
func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("[%s] %d\n", label, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("===== GOROUTINES =====")

	// 1. WITHOUT goroutines (Sequential)
	// Each function runs to completion before the next starts.
	fmt.Println("\n--- Sequential (No Goroutines) ---")
	start := time.Now()
	printNumbers("A")
	printNumbers("B")
	fmt.Printf("Sequential took: %v\n", time.Since(start))

	// 2. WITH goroutines (Concurrent)
	// Both functions run at the SAME time!
	fmt.Println("\n--- Concurrent (With Goroutines) ---")
	start = time.Now()
	go printNumbers("A") // Launches goroutine (runs in background)
	go printNumbers("B") // Launches another goroutine
	// Without waiting, main() would exit and kill both goroutines!
	time.Sleep(600 * time.Millisecond) // Ugly but simple wait for now
	fmt.Printf("Concurrent took: %v\n", time.Since(start))

	// 3. Anonymous Goroutine
	fmt.Println("\n--- Anonymous Goroutine ---")
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from an anonymous goroutine!")
	time.Sleep(50 * time.Millisecond)

	// 4. The Problem: Main exits too early!
	// If main() finishes, ALL goroutines are killed instantly.
	// That's why we used time.Sleep above — but it's unreliable.
	// The proper solution? Channels or WaitGroups (next lesson!)
	fmt.Println("\n⚠️  Using time.Sleep to wait is BAD. Next: Channels!")
}
