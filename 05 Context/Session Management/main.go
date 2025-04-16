package main

import (
	"context"
	"fmt"
	"session-management/session"
)

// This program manages session-related data using Go's context package.
// It allows storing and retrieving session information such as:
// User ID, session ID, and admin access rights within the context.
// This is useful in web applications where user and session-specific data need to be passed along the call chain without using global variables.

func main() {
	// Create a background context (starting point for session data)
	ctx := context.Background()

	// Set session data in the context
	ctx = session.SetUserID(ctx, 55)
	ctx = session.SetSessionID(ctx, "1536")
	ctx = session.SetAdminAccess(ctx, true)

	// Retrieve session data from the context
	userID := session.GetUserID(ctx)
	sessionID := session.GetSessionID(ctx)
	isAdmin := session.GetAdmin(ctx)

	// Print the retrieved session information
	fmt.Println("id:", userID)
	fmt.Println("admin:", isAdmin)
	fmt.Println("session:", sessionID)
}
