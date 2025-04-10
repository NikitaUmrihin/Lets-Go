package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This program demonstrates the basic use of channels.
// Channels in Go are used for safe communication between goroutines
// They can be used to send and receive data in a synchronized manner

// produce generates 10 random integers and sends them into the channel
func produce(intChannel chan int) {
	defer close(intChannel)
	for i := 0; i < 10; i++ {
		intChannel <- rand.Intn(100)
	}
}

// consume receives and prints integers from the channel until it's closed
func consume(intChannel chan int) {
	for num := range intChannel {
		fmt.Println("Received:", num)
		time.Sleep(time.Millisecond * 100)

	}
}

func main() {

	// Create an unbuffered integer channel
	intChannel := make(chan int)

	// Start production separate goroutine
	go produce(intChannel)

	// Consume numbers in the main goroutine
	consume(intChannel)
}
