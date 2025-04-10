package main

import (
	"fmt"
	"sync"
)

// This program demonstrates the how to use of WaitGroup.
// WaitGroup is used to wait for multiple Goroutines to finish execution before the main function proceeds.
// The 4 steps are:
//   - Add tasks → wg.Add(n):
//     Increases the counter by number of Goroutines
//   - Run Goroutines → go func():
//     Each Goroutine does its work
//   - Mark completion → wg.Done()
//     Decreases the counter by 1 when a Goroutine finishes
//   - Wait for all → wg.Wait()
//     Blocks execution until the counter reaches 0
//________________________________________________________________________________________________________

// sumSlice computes the sum of a slice of integers and sends the result to a channel
// The function takes:
//   - nums []int: A slice of integers.
//   - wg *sync.WaitGroup: A pointer to a WaitGroup, used to wait for Goroutines to complete.
//   - ch chan int: A channel for sending computed sums back to the main function.
func sumSlice(nums []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Mark this function as complete when it finishes execution
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum
}

// The computeSum function takes a slice of integers (numbers) and a specified number of partitions (n).
// It splits the slice into 'n' parts, computes the sum of each part concurrently using Goroutines, and returns the total sum.
func computeSum(numbers []int, n int) int {

	if n > len(numbers) {
		fmt.Println("Warning: Number of partitions is greater than the array size.")
		fmt.Println("Exiting without computation.")
		return 0
	}

	size := (len(numbers) + n - 1) / n // Compute chunk size

	// Declare a WaitGroup to track Goroutines
	var wg sync.WaitGroup

	// Create a buffered channel to store the sums
	ch := make(chan int, n)

	// Split array and launch goroutines to compute sum concurrently
	for i := 0; i < len(numbers); i += size {
		// Define the end index of the current chunk
		end := i + size

		// Make sure we don’t go out of bounds
		if end > len(numbers) {
			end = len(numbers)
		}

		// Increase the WaitGroup counter before starting a Goroutine
		wg.Add(1)

		// Start a Goroutine to compute the sum
		go sumSlice(numbers[i:end], &wg, ch)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	close(ch)

	// In Go, reading from a closed channel is allowed !
	// When a channel is closed, we can still receive all values that were sent before closing.

	totalSum := 0

	// When using a for range loop over the channel
	// THE CHANNEL MUST BE ALREADY CLOSED !

	// Read values incoming from channel
	for partialSum := range ch {
		// Combine results
		totalSum += partialSum
	}

	return totalSum
}

//
// 	MAIN :)
//

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	parts := 11 // Number of partitions

	totalSum := computeSum(numbers, parts)
	fmt.Println("Total Sum:", totalSum)
}
