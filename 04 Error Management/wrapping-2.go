package main

import (
	"errors"
	"fmt"
)

// This program demonstrates two ways of how to wrap and unwrap errors in Go.
// The programs functions, 'goToUrl' and 'searchUrl' simulate URL-related errors.
// The errors are then wrapped and checked using Go’s built-in errors.Is method.

// _________________________________________________________________________

// UrlNotFoundError is a custom error variable
var UrlNotFoundError = fmt.Errorf("NOT FOUND")

// _________________________________________________________________________

// Define wrapper error type 'notOkError'
// which adds additional context (status) to a base error.

type notOkError struct {
	Url    string // Url that created the error
	Status int    // Status code associated with the error
	base   error  // The base error being wrapped
}

// Implement error interface
func (e notOkError) Error() string {
	return fmt.Sprintf("%s: %d %s", e.Url, e.Status, e.base)
}

// Implement Unwrap() so we can use Is()
func (e notOkError) Unwrap() error {
	return e.base
}

// _________________________________________________________________________

// Functions to simulate error situations

// goToUrl returns a "404 NOT FOUND" error
func goToUrl(url string) (string, error) {
	return "", UrlNotFoundError
}

// searchUrl returns a wrapped 'notOkError' with the 'UrlNotFoundError' as the base
func searchUrl(url string) (string, error) {
	return "", notOkError{
		Url:    url,
		Status: 404,
		base:   UrlNotFoundError,
	}
}

// _________________________________________________________________________

func main() {
	url := "google.com"

	// Testing the first error wrapping method
	// 	- Using fmt.Errorf with %w to wrap an error
	fmt.Println("Testing gotToUrl() - 1st way of wrapping errors -> using Errorf()")

	var testNotOkErr notOkError

	_, err := goToUrl(url)

	if err != nil {
		// Wrap the error with the URL in the error message
		wrapped := fmt.Errorf("%s: %w", url, err)

		// Is() and As() will only work if Unwrap() is implemented !!!

		// Check if the wrapped error is of the type UrlNotFoundError
		if errors.Is(wrapped, UrlNotFoundError) {
			fmt.Println("Yes! This is a UrlNotFoundError")
		}

		// With custom error struct we cannot use Is() - unless we're comparing the same instance
		if errors.As(err, &testNotOkErr) {
			fmt.Println("Yes! This is a notOkError")
		}

		fmt.Println(wrapped)
	}

	// Testing the second error wrapping method
	// 	- Using a custom error type
	fmt.Println("\nTesting searchUrl() - 2nd way of wrapping errors -> using wrapper struct")

	_, err = searchUrl(url)

	if err != nil {

		// Check if the unwrapped error is of the type UrlNotFoundError
		if errors.Is(err, UrlNotFoundError) {
			fmt.Println("Yes! This is a UrlNotFoundError")
		}

		// With custom error struct we cannot use Is() - unless we're comparing the same instance
		if errors.As(err, &testNotOkErr) {
			fmt.Println("Yes! This is a notOkError")
		}

		fmt.Println(err)
	}

}
