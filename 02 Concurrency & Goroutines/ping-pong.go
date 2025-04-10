package main

import (
	"fmt"
	"time"
)

//  This program demonstrates concurrent communication between two goroutines
//  using a channel to simulate a ping pong message exchange.

// pinger continuously waits for a message from the receiver channel,
// prints "PING", sleeps and then sends a message to the transmitter channel.
func pinger(channel chan string) {
	for {
		// Wait to receive a message
		<-channel

		// Once received - ping the ponger
		fmt.Println("\tPING")
		time.Sleep(time.Millisecond * 250)
		channel <- "ping"
	}
}

// ponger behaves similarly to pinger but prints "PONG" instead.
// It listens on its receiver channel, sleeps, and then transmits to the other channel.
func ponger(channel chan string) {
	for {
		// Wait to receive a message
		<-channel

		// Once received - pong the pinger
		fmt.Println("\t\t\t\t\t\tPONG")
		time.Sleep(time.Millisecond * 250)
		channel <- "pong"
	}
}

func main() {

	// Channel used for ping-pong communication between goroutines.
	channel := make(chan string)

	// Start goroutines
	go pinger(channel)
	go ponger(channel)

	// Kickstart the ping pong cycle
	channel <- "Let's Go"

	// Keep the main function alive so goroutines can continue running.
	fmt.Scanln()
}
