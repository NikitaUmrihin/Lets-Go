package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// The program launches 50 goroutines, each performing simulated work and printing messages.
// It uses a cancelable context to manage their lifecycle.
// After a short delay, it prints the active goroutine count, cancels the context to stop them and prints the final count.

const WORK_DURATION = 50
const WAIT_DURATION = 100

func main() {
	fmt.Printf("GOROUTINES RUNNING %d\n", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelF := context.WithCancel(ctx)
	defer cancelF()

	// Start 50 goroutines
	for i := range 50 {
		// Inline (anonymous) functions in Go:
		// Must always be declared using the func keyword (without a name)
		go func(n int) {
			fmt.Println("STARTING GOROUTINE #", n)
			for {
				select {
				case <-ctx.Done():
					// If the context is canceled, exit the goroutine
					runtime.Goexit()
					// return
				default:
					// Simulate work by printing and sleeping
					fmt.Printf("routine %d still working\n", n)
					time.Sleep(WORK_DURATION * time.Millisecond)
				}
			}
		}(i)
	}
	// Allow some time for goroutines to start before checking count
	time.Sleep(time.Millisecond)
	fmt.Printf("\nGOROUTINES RUNNING: %d\n\n", runtime.NumGoroutine())

	// Cancel the context, wait to ensure goroutines have time to terminate
	cancelF()
	time.Sleep(WAIT_DURATION * time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING: %d (AFTER CANCEL FUNCTION CALLED)\n", runtime.NumGoroutine())

}
