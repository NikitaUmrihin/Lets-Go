package main

import (
	"errors"
	"fmt"
)

// This program demonstrates how to access wrapped errors in Go using errors.As.
// It wraps a UrlNotFoundError in a custom error type (notOkError)
// And uses errors.As to extract and print the status code from the wrapped error.

// _________________________________________________________________________

// Define a custom error type 'notOkError' which adds additional context (status) to a base error.
type notOkError struct {
	status int
	Base   error
}

// Implement error interface
func (e notOkError) Error() string {
	return fmt.Sprintf("%d %s", e.status, e.Base)
}

func (e notOkError) Unwrap() error {
	return e.Base
}

// _________________________________________________________________________

// Custom error variable
var UrlNotFoundError = fmt.Errorf("URL NOT FOUND")

// _________________________________________________________________________

// Functions to simulate error situations

// openUrl returns a wrapped 'notOkError' with the "404 NOT FOUND" error as the base
func openUrl(url string) (string, error) {
	return "", notOkError{
		status: 404,
		Base:   UrlNotFoundError,
	}
}

// searchUrl calls openUrl and wraps the error with additional context.
func searchUrl(url string) error {
	_, err := openUrl(url)
	if err != nil {
		// Wrap the error with a custom message
		return fmt.Errorf("Error while searching: %s %w", url, err)
	}
	// Search algorithm....
	return nil
}

// _________________________________________________________________________

func main() {
	url := "google.com"

	err := searchUrl(url)

	if err != nil {

		// Try to access the original 'notOkError' using 'errors.As'
		var urlErr notOkError

		// Check if error is 'notOkError', and handle error according to its status
		if errors.As(err, &urlErr) {

			if urlErr.status == 400 {
				fmt.Println("This looks like a Bad Request...")
			}

			if urlErr.status == 404 {
				fmt.Println("Please double check the URL")
			}

			if urlErr.status == 429 {
				fmt.Println("Too many requests...\nTry again later")
			}

			// Print error status code
			fmt.Printf("ERROR CODE %d\n", urlErr.status)
			fmt.Println("Yes! This is a notOkError")
		}

		// Print the wrapped error
		fmt.Println(err)
	}

}
