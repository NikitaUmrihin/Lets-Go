package main

import (
	"errors"
	"fmt"
)

// This Go program demonstrates why custom errors should not be compared directly.
// It shows that custom errors are distinct (even with the same message),
// while predefined errors are comparable.

// _________________________________________________________________________

// FileNotFound creates a custom error
func FileNotFound(path string) error {

	// Using fmt.Errorf to format and return a custom error with the file path
	return fmt.Errorf("File does not exist %d", path)
}

// _________________________________________________________________________

// UserNotFound is a custom error variable
var UserNotFound = errors.New("User does not exist")

// _________________________________________________________________________
func main() {
	// Defining two file names for testing
	file := "1.txt"

	//	Create 2 FileNotFound errors (with same file)
	fErr1 := FileNotFound(file)
	fErr2 := FileNotFound(file)

	// Compare errors:
	// Even though they are both FileNotFound
	fmt.Println(fErr1 == fErr2) // The statement will return false

	//	Create 2 UserNotFound errors
	uErr1 := UserNotFound
	uErr2 := UserNotFound

	// Compare errors: will return true (it's the same predefined error)
	fmt.Println(uErr1 == uErr2)
}
