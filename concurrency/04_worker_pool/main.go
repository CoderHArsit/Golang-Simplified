package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================================
// WORKER POOL PATTERN
// ============================================================
// A Worker Pool limits the number of goroutines processing tasks.
//
// Without a pool: 10,000 tasks → 10,000 goroutines (can overwhelm the system)
// With a pool:    10,000 tasks → 5 workers handle them one by one
//
// Architecture:
//   [Producer] → jobs channel → [Worker 1]
//                              [Worker 2]  → results channel → [Collector]
//                              [Worker 3]
// ============================================================

// Job represents a unit of work
type Job struct {
	ID       int
	Data     int
}

// Result represents a completed job
type Result struct {
	Job    Job
	Output int
}

// worker reads jobs from the 'jobs' channel and writes results to the 'results' channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate variable processing time
		duration := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(duration)

		output := job.Data * job.Data // Square the number
		results <- Result{Job: job, Output: output}
		fmt.Printf("  Worker %d: processed Job %d (data=%d → result=%d) in %v\n",
			id, job.ID, job.Data, output, duration)
	}
}

func main() {
	fmt.Println("===== WORKER POOL =====")

	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 1. Launch Workers
	// Workers start and immediately block on the 'jobs' channel, waiting for work.
	var wg sync.WaitGroup
	fmt.Printf("Starting %d workers...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Send Jobs
	// The producer sends jobs into the channel.
	fmt.Printf("Sending %d jobs...\n\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: rand.Intn(50)}
	}
	close(jobs) // No more jobs → workers will exit their 'range' loop

	// 3. Wait for Workers to Finish, Then Close Results
	go func() {
		wg.Wait()
		close(results) // Safe to close now that all workers are done
	}()

	// 4. Collect Results
	fmt.Println("\n--- Collected Results ---")
	totalResults := 0
	for r := range results {
		fmt.Printf("Job %d: %d² = %d\n", r.Job.ID, r.Job.Data, r.Output)
		totalResults++
	}
	fmt.Printf("\nProcessed %d jobs with %d workers.\n", totalResults, numWorkers)
}
