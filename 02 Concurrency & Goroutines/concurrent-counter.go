package main

import (
	"fmt"
	"sync"
)

// This program demonstrates the basic use of Mutex (and WaitGroup).
// Mutex in Go are used to control access to shared data,
// so only one goroutine can use it at a time.

// Declare a shared counter variable
var counter int

// Create a mutex to control access to the counter
var mu sync.Mutex

// increment safely increases the counter using a mutex
// It takes the goroutine number and a WaitGroup pointer
func increment(num int, wg *sync.WaitGroup) {

	defer wg.Done()

	// Lock() ensures only one can modify the counter at a time
	mu.Lock()
	fmt.Println("Go routine:", num, " -> counter:", counter)
	counter++
	// Unlock the mutex so other goroutines can proceed
	mu.Unlock()
}

// main launches 10 goroutines that concurrently increment a shared counter
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(i, &wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
