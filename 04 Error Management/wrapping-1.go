package main

import (
	"errors"
	"fmt"
)

// This program demonstrates error wrapping and unwrapping in Go.
// It creates a chain of wrapped errors using Errorf() with the %w verb,
// then progressively unwraps them using Unwrap() to access the original error.

func makeError() error {
	return errors.New("Error #1")
}

func makeError2() error {
	return fmt.Errorf("Error #2 wraps (%w)", makeError())
}

func makeError3() error {
	return fmt.Errorf("Error #3 wraps [ %w ]", makeError2())
}

func makeError4() error {
	return fmt.Errorf("Error #4 wraps { %w }", makeError3())
}

// main shows how to build and traverse an error chain step by step.
func main() {
	err := makeError4()
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)
}
