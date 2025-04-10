package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This program demonstrates the basic use of select keyword (just like switch -> case).
// It simulates concurrent data from two cats messaging each other on chat app.

// keyboard contains the characters used to generate random messages by cats
var keyboard = []rune("abcdefghijklmnopqrstuvwxyz /*-+0123456789.=][;'.,ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// blackCat simulates a Black Cat on a keyboard sending messages at random intervals.
func blackCat(ch chan string) {
	for {
		// Sleep randomly for 1-2 seconds (to mimic sensor delay)
		sleepTime := time.Duration(1+rand.Intn(2)) * time.Second
		time.Sleep(sleepTime)

		// Generate a random string
		msg := make([]rune, rand.Intn(16)+5)
		for i := range msg {
			msg[i] = keyboard[rand.Intn(len(keyboard))]
		}

		ch <- string(msg)
	}
}

// blackCat simulates an Orange Cat on a keyboard sending messages at random intervals.
func orangeCat(ch chan string) {
	for {
		sleepTime := time.Duration(1+rand.Intn(2)) * time.Second
		time.Sleep(sleepTime)

		// Generate a random string
		msg := make([]rune, rand.Intn(25)+8)
		for i := range msg {
			msg[i] = keyboard[rand.Intn(len(keyboard))]
		}

		ch <- string(msg)
	}
}

// Let's Go!

func main() {

	// Create channels for users on chat app
	black := make(chan string)
	orange := make(chan string)

	// Launch goroutines
	go blackCat(black)
	go orangeCat(orange)

	// Set a timeout to limit the chat duration
	monitorDuration := 10 * time.Second
	timeout := time.After(monitorDuration)
	fmt.Printf("Starting chat conversation %v...\n", monitorDuration)

	// Use select in a loop to handle messages as soon they are available
	for {

		// 'select {case}' - is just like 'switch {case}' , but for channels !
		select {

		case temp := <-black:
			fmt.Println("User 1: ", temp)

		case hum := <-orange:
			fmt.Println("User 2: ", hum)

		// If the timeout channel signals (after 10 seconds), exit the loop
		case <-timeout:
			fmt.Println("Conversation ended: timeout reached.")
			return
		}
	}
}
