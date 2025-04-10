package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main continuously reads user input from the terminal and appends each entry to a log file.
// The program exits when the user types "exit".
func main() {
	// Define the log file name
	logFile := "log.txt"

	// Open the log file in append mode, create it if it doesn't exist, and ensure it's writable
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}

	// Ensure the file is closed after the program finishes
	defer file.Close()

	// Create a new scanner to read input from the keyboard
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Log Writer (type 'exit' to quit)")

	// Loop to continuously read user input
	for {
		fmt.Print("Enter log message: ")

		// If there's no more input, break the loop
		if !scanner.Scan() {
			break
		}

		// Get the text entered by the user
		text := scanner.Text()

		// If the user types "exit", break the loop
		if strings.ToLower(text) == "exit" {
			fmt.Println("Exiting log writer.")
			break
		}

		// Write the log message to the file
		_, err := file.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Error writing to log file:", err)
			break
		}

		fmt.Println("Log entry saved.")
	}

	// If there's an error reading input, print error message
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
