package main

import (
	"context"
	"fmt"
	"time"
)

// This Go program demonstrates how to use a context with a timeout.
// It creates a context that times out after 500 milliseconds and simulates work with a 100ms sleep.
// It then checks if the context has finished (either due to a timeout or cancellation).
// The program helps manage timeouts in operations, ensuring they don't run indefinitely.

// TIMEOUT_DURATION > WORK_DURATION = no timeout
// TIMEOUT_DURATION < WORK_DURATION = timeout
const TIMEOUT_DURATION = 500
const WORK_DURATION = 100

func main() {
	// Create a new context with a timeout (500 milliseconds)
	ctx := context.Background()
	ctx, cancelF := context.WithTimeout(ctx, TIMEOUT_DURATION*time.Millisecond)
	defer cancelF() // Ensure that the cancel function is called when main finishes

	// Simulate work by making the program sleep for 100 milliseconds
	time.Sleep(WORK_DURATION * time.Millisecond)

	// The select statement listens for the context to be done (timed out or canceled)
	select {
	// If the context is done
	case <-ctx.Done():
		fmt.Println("didnt finish in time")
	default:
		fmt.Println("finished")
	}
}
