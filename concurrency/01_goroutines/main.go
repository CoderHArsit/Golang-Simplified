package main

import (
	"fmt"
	"runtime"
	"time"
)

// ============================================================
// CONCURRENCY vs PARALLELISM (Interview Classic!)
// ============================================================
//
// Concurrency = DESIGN — structuring your program to handle
//   multiple things. It's about DEALING with many things at once.
//
// Parallelism = EXECUTION — physically running multiple things
//   at the same time on multiple CPU cores. It's about DOING
//   many things at once.
//
// Analogy:
//   Concurrency: One chef switching between 3 dishes on the stove.
//   Parallelism:  Three chefs each cooking one dish simultaneously.
//
// Go is concurrent by default. It becomes parallel when:
//   - runtime.GOMAXPROCS > 1 (default = NumCPU since Go 1.5)
//   - Goroutines are scheduled across multiple OS threads
//
// Key insight: You can have concurrency without parallelism
//   (single-core CPU), but you can't have parallelism without
//   concurrency (you need the structure first).
// ============================================================

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
	// We define a function inline and launch it as a goroutine.
	// Notice: we PASS the string as an argument instead of capturing it from outer scope.
	fmt.Println("\n--- Anonymous Goroutine ---")
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from an anonymous goroutine!")
	time.Sleep(50 * time.Millisecond)

	// 4. WHY pass arguments instead of using closures?
	//
	// A closure captures variables by REFERENCE, not by value.
	// If the variable changes before the goroutine runs, you get unexpected results.
	//
	// Example (closure — risky):
	//   msg := "Hello"
	//   go func() {
	//       fmt.Println(msg)  // ← captures `msg` by reference
	//   }()
	//   msg = "Goodbye"       // goroutine might print "Goodbye"!
	//
	// Example (parameter — safe):
	//   msg := "Hello"
	//   go func(m string) {
	//       fmt.Println(m)    // ← `m` is a local copy, safe!
	//   }(msg)
	//   msg = "Goodbye"       // goroutine still prints "Hello" ✅

	// 5. The Classic Loop Trap 🪤
	// This is where it REALLY matters. Without passing `i` as a param,
	// all goroutines share the same loop variable and likely print "5".
	fmt.Println("\n--- Loop Trap: Closure (BAD) ---")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("  closure i = %d\n", i) // ❌ All might print 5!
		}()
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n--- Loop Trap: Parameter (GOOD) ---")
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("  param n = %d\n", n) // ✅ Prints 0,1,2,3,4 (any order)
		}(i) // `i` is copied into `n` at launch time
	}
	time.Sleep(100 * time.Millisecond)

	// 6. The Problem: Main exits too early!
	// If main() finishes, ALL goroutines are killed instantly.
	// That's why we used time.Sleep above — but it's unreliable.
	// The proper solution? Channels or WaitGroups (next lesson!)

	// 7. runtime.GOMAXPROCS — Controlling Parallelism
	// GOMAXPROCS sets how many OS threads can run goroutines simultaneously.
	// Default = number of CPU cores (since Go 1.5).
	fmt.Println("\n--- GOMAXPROCS ---")
	fmt.Printf("  CPU cores available: %d\n", runtime.NumCPU())
	fmt.Printf("  GOMAXPROCS (current): %d\n", runtime.GOMAXPROCS(0)) // 0 = query, don't change
	fmt.Printf("  Goroutines running: %d\n", runtime.NumGoroutine())
	// You CAN set it: runtime.GOMAXPROCS(1) — forces single-threaded execution.
	// But you almost NEVER need to change it in production.

	fmt.Println("\n⚠️  Using time.Sleep to wait is BAD. Next: Channels!")
}
