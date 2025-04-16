package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// This program demonstrates making an HTTP GET request with a timeout using Go's context package.
// It sends a request to a specified URL, checks if the request is completed within the context's deadline.
// The program prints the response if successful or raises an error if it encounters issues.

const DEADLINE = 2500

// makeRequest sends an HTTP GET request to the specified route with a given context.
func makeRequest(ctx context.Context, route string) (string, error) {

	// Check if the context has a deadline and if it's close to being exceeded
	deadline, ok := ctx.Deadline()

	// If the deadline is too close, return an error
	if ok && time.Until(deadline) < 100*time.Millisecond {
		return "", fmt.Errorf("Deadline is coming too soon")
	}

	// Create a new HTTP request using the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, route, nil)
	if err != nil {
		return "", err
	}

	// Perform the HTTP request using the default HTTP client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// Ensure the response body is closed after the function exits
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	// Read the response body into a byte slice
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Return the response body as a string
	return string(bs), nil
}

// main function demonstrates the use of a timeout context when making an HTTP request
func main() {
	// Create a context with a timeout of 25 milliseconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*DEADLINE)
	defer cancel()

	// Attempt to make a request to google.com
	resp, err := makeRequest(ctx, "https://youtube.com")

	// If there's an error (e.g., timeout), panic and stop the program
	if err != nil {
		panic(err)
	}

	// Print the response body if the request succeeds
	fmt.Println(resp)
}
