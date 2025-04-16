package main

import (
	"context"
	"fmt"
)

// This program demonstrates how to use Go's context package to store and retrieve a user ID.
// It checks if the user ID exists in the context and prints it, or prints "Not logged in" if not found.

// Define a custom key type for the context to avoid key collisions
type key string

// Create a key to store the user ID in the context
var userKey key = "UserID"

// With Value Context Example
func main() {
	// Create a context with a value (user ID) and attach it to the background context
	ctx := context.WithValue(context.Background(), userKey, 1)

	// Retrieve the value stored in the context for the userKey
	userId, ok := ctx.Value(userKey).(int)

	// If the value is not of type int or not found, print "Not logged in!"
	if !ok {
		fmt.Println("Not logged in!")
		return
	}
	// Print the retrieved user ID
	fmt.Println(userId)
}
