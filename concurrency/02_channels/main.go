package main

import "fmt"

func main() {
	fmt.Println("===== CHANNELS =====")

	// 1. Basic Unbuffered Channel
	// An unbuffered channel BLOCKS the sender until the receiver is ready.
	// Think of it as a walkie-talkie: you can't talk until the other person is listening.
	fmt.Println("\n--- Unbuffered Channel ---")
	ch := make(chan string) // Create a channel that carries strings

	go func() {
		ch <- "Hello from goroutine!" // SEND into channel (blocks until someone reads)
	}()

	msg := <-ch // RECEIVE from channel (blocks until someone sends)
	fmt.Println("Received:", msg)

	// 2. Channels as Synchronization
	// Instead of time.Sleep, we use a channel to WAIT for a goroutine to finish.
	fmt.Println("\n--- Channel as Signal ---")
	done := make(chan bool)

	go func() {
		fmt.Println("  Worker: doing heavy work...")
		// Simulate work
		for i := 0; i < 3; i++ {
			fmt.Printf("  Worker: step %d\n", i+1)
		}
		done <- true // Signal: "I'm done!"
	}()

	<-done // Block until we receive the signal
	fmt.Println("Main: worker finished!")

	// 3. Buffered Channels
	// A buffered channel allows N sends before blocking.
	// Think of it as a mailbox that can hold N letters.
	fmt.Println("\n--- Buffered Channel ---")
	mailbox := make(chan string, 3) // Buffer size = 3

	mailbox <- "Letter 1" // Does NOT block (buffer has room)
	mailbox <- "Letter 2"
	mailbox <- "Letter 3"
	// mailbox <- "Letter 4" // THIS WOULD BLOCK! Buffer is full.

	fmt.Println(<-mailbox) // "Letter 1" (FIFO)
	fmt.Println(<-mailbox) // "Letter 2"
	fmt.Println(<-mailbox) // "Letter 3"

	// 4. Ranging over a Channel
	// Use 'range' to read from a channel until it's CLOSED.
	fmt.Println("\n--- Range over Channel ---")
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers) // MUST close, or the range will wait forever (deadlock!)
	}()

	for n := range numbers {
		fmt.Printf("Received: %d\n", n)
	}

	// 5. Directional Channels
	// You can restrict a channel to send-only or receive-only in function signatures.
	fmt.Println("\n--- Directional Channels ---")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Hello!")
	pong(pings, pongs)
	fmt.Println("Pong received:", <-pongs)

	// 6. Channel Axioms — nil & closed channel behavior
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
	//   - Closing is only necessary when the receiver needs to know
	//     that no more values are coming (e.g., range loop).
	//   - You CAN read from a closed channel (returns zero values).
	//   - You CANNOT send to a closed channel (panics!).
	//   - nil channels are useful in select to disable a case.
	// ============================================================
	fmt.Println("\n--- Channel Axioms ---")

	// Reading from a CLOSED channel returns zero value + false
	axiomCh := make(chan int, 2)
	axiomCh <- 42
	axiomCh <- 99
	close(axiomCh)

	val1, ok1 := <-axiomCh
	fmt.Printf("Read 1: value=%d, ok=%t (buffered data)\n", val1, ok1)
	val2, ok2 := <-axiomCh
	fmt.Printf("Read 2: value=%d, ok=%t (buffered data)\n", val2, ok2)
	val3, ok3 := <-axiomCh
	fmt.Printf("Read 3: value=%d, ok=%t ← closed! zero value returned\n", val3, ok3)

	// The comma-ok idiom: always check 'ok' to detect closed channels
	// if !ok { fmt.Println("Channel closed, stop reading!") }
	//
	// PANIC examples (DO NOT uncomment!):
	//   axiomCh <- 1       // panic: send on closed channel
	//   close(axiomCh)     // panic: close of closed channel
	//
	//   var nilCh chan int
	//   close(nilCh)       // panic: close of nil channel
}

// send-only channel (chan<-)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// receive from one, send to another
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
