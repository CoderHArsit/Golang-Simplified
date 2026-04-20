package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("===== SYNC EXTRAS =====")

	// ============================================================
	// 1. sync.Once — Exactly-Once Initialization
	// ============================================================
	// sync.Once ensures a function runs EXACTLY ONCE, even when
	// called from multiple goroutines simultaneously.
	//
	// Common use cases:
	//   - Singleton pattern (database connection, logger)
	//   - Lazy initialization
	//   - One-time configuration loading
	//
	// Interview Q: "How do you implement a singleton in Go?"
	// Answer: Use sync.Once — it's goroutine-safe and idiomatic.
	fmt.Println("\n--- sync.Once ---")

	var once sync.Once
	var wg sync.WaitGroup

	initDB := func() {
		fmt.Println("  Initializing database connection... (runs only once!)")
	}

	// Launch 5 goroutines — ALL try to initialize, only ONE succeeds
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d: trying to init DB\n", id)
			once.Do(initDB) // Only the first call executes initDB
			fmt.Printf("  Goroutine %d: DB is ready\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("  All goroutines done. DB was initialized exactly once. ✅")

	// ============================================================
	// 2. sync.RWMutex — Reader/Writer Lock
	// ============================================================
	// sync.RWMutex allows:
	//   - MULTIPLE concurrent readers (RLock/RUnlock)
	//   - Only ONE writer at a time (Lock/Unlock)
	//   - Writers block all readers AND other writers
	//
	// When to use: Read-heavy workloads (caches, configs, maps)
	//   - If reads >> writes → RWMutex is faster than Mutex
	//   - If reads ≈ writes → Just use regular Mutex
	//
	// Interview Q: "Difference between Mutex and RWMutex?"
	// Answer: RWMutex allows concurrent reads; Mutex does not.
	fmt.Println("\n--- sync.RWMutex ---")

	type SafeCache struct {
		mu   sync.RWMutex
		data map[string]string
	}

	cache := SafeCache{data: make(map[string]string)}

	// Writer: needs exclusive access
	write := func(key, value string) {
		cache.mu.Lock() // Exclusive lock — blocks readers AND writers
		defer cache.mu.Unlock()
		cache.data[key] = value
		fmt.Printf("  WRITE: %s = %s\n", key, value)
	}

	// Reader: can run concurrently with other readers
	read := func(key string) string {
		cache.mu.RLock() // Shared lock — allows other readers
		defer cache.mu.RUnlock()
		return cache.data[key]
	}

	// Demo: Write some values, then read concurrently
	write("name", "Harshit")
	write("lang", "Go")

	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			val := read("name")
			fmt.Printf("  Reader %d: name = %s\n", id, val)
		}(i)
	}
	wg2.Wait()

	// ============================================================
	// 3. sync.Map — Concurrent-Safe Map
	// ============================================================
	// Regular Go maps are NOT goroutine-safe. Concurrent read+write
	// causes: "fatal error: concurrent map writes"
	//
	// Options:
	//   a) map + sync.Mutex (most common, fine for most cases)
	//   b) sync.Map (optimized for specific patterns)
	//
	// sync.Map is best when:
	//   - Keys are stable (written once, read many times)
	//   - Many goroutines read/write disjoint sets of keys
	//
	// Use map + Mutex when:
	//   - You need to iterate frequently
	//   - You know all keys upfront
	//   - You need type safety (sync.Map uses interface{})
	//
	// Interview Q: "Is Go's map goroutine-safe?"
	// Answer: NO. Use sync.Map or a map protected by sync.Mutex.
	fmt.Println("\n--- sync.Map ---")

	var sm sync.Map

	// Store values (like map[key] = value)
	var wg3 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			key := fmt.Sprintf("worker_%d", id)
			sm.Store(key, id*10)
			fmt.Printf("  Stored: %s = %d\n", key, id*10)
		}(i)
	}
	wg3.Wait()

	// Load a value
	if val, ok := sm.Load("worker_3"); ok {
		fmt.Printf("  Loaded: worker_3 = %v\n", val)
	}

	// Range over all entries (like for k, v := range m)
	fmt.Println("  All entries:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("    %v → %v\n", key, value)
		return true // return false to stop iteration
	})

	// ============================================================
	// 4. errgroup — Error-Aware Concurrency (Production Go)
	// ============================================================
	// errgroup is from golang.org/x/sync — used in production MORE
	// than raw WaitGroups because it handles errors properly.
	//
	// WaitGroup: Waits for goroutines, ignores errors.
	// errgroup:  Waits for goroutines AND returns the FIRST error.
	//
	// Install: go get golang.org/x/sync
	//
	// Interview Q: "How do you handle errors in concurrent Go code?"
	// Answer: Use errgroup. It's cleaner than manually collecting
	// errors through channels.
	fmt.Println("\n--- errgroup ---")

	g := new(errgroup.Group)

	// Launch concurrent tasks
	urls := []string{"google.com", "github.com", "INVALID_URL"}

	for _, url := range urls {
		url := url // Capture for goroutine (or pass as param)
		g.Go(func() error {
			return fetchURL(url)
		})
	}

	// Wait for ALL goroutines and get the FIRST error
	if err := g.Wait(); err != nil {
		fmt.Printf("  Error: %v ❌\n", err)
	} else {
		fmt.Println("  All fetches succeeded! ✅")
	}

	// errgroup with CONCURRENCY LIMIT (SetLimit)
	fmt.Println("\n--- errgroup with Limit ---")
	g2 := new(errgroup.Group)
	g2.SetLimit(2) // Max 2 concurrent goroutines (like a semaphore!)

	for i := 1; i <= 5; i++ {
		i := i
		g2.Go(func() error {
			fmt.Printf("  Task %d: running (max 2 concurrent)\n", i)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("  Task %d: done\n", i)
			return nil
		})
	}
	g2.Wait()
	fmt.Println("  All limited tasks completed! ✅")

	// ============================================================
	// SUMMARY
	// ============================================================
	// sync.Once    → One-time init, singleton pattern
	// sync.RWMutex → Multiple readers OR one writer
	// sync.Map     → Concurrent-safe map (use sparingly)
	// errgroup     → WaitGroup + error handling (production Go)
	fmt.Println("\n✅ Sync extras complete!")
}

// Simulates fetching a URL with possible failure
func fetchURL(url string) error {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if url == "INVALID_URL" {
		return fmt.Errorf("failed to fetch %s", url)
	}
	fmt.Printf("  Fetched: %s ✅\n", url)
	return nil
}
