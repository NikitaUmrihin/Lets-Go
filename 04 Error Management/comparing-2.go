package main

import (
	"errors"
	"fmt"
	"time"
)

// This Go program demonstrates how to compare custom errors using the errors.Is method.
// It defines a custom FileNotFoundError type and implements the Is method to check if an error is of the same type, ignoring its field values

// _________________________________________________________________________

// Define a custom error type 'FileNotFoundError'
type FileNotFoundError struct {
	Filename string
	When     time.Time
}

// Implement error interface
func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("File %s doesn't exist\n%v", e.Filename, e.When)
}

// Implement Is() to check if the Error Is the same type (ignoring fields)
func (e FileNotFoundError) Is(other error) bool {
	_, ok := other.(FileNotFoundError)
	return ok
}

// _________________________________________________________________________

func main() {

	// Create an instance of FileNotFoundError error
	err1 := FileNotFoundError{
		Filename: "test.txt",
		When:     time.Now(),
	}

	err2 := FileNotFoundError{
		Filename: "best.jpg",
		When:     time.Now(),
	}

	// Check if the created errors are of the same type (FileNotFoundError)
	isTheSameError := errors.Is(err1, err2)
	fmt.Println(isTheSameError)
}
