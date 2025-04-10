package main

import (
	"fmt"
	"time"
)

// This program demonstrates the basic use of Goroutines.

// letsGo prints a message 20 times with a 500ms delay between each
func letsGo(st string) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d) %s\n", i, st)
		time.Sleep(time.Millisecond * 500)
	}
}

// Implicit goroutine - thread 1
func main() {
	go letsGo("eating breakfast") // Explicit goroutine - thread 2
	go letsGo("drinking coffee")  // Explicit goroutine - thread 3
	go letsGo("watching TV")

	// Block the main goroutine until user input is received
	// This keeps the program alive long enough to observe goroutines running
	fmt.Scanln()

	// Alternative: use time.Sleep to let goroutines finish
	// time.Sleep(time.Millisecond * 3000)
}
