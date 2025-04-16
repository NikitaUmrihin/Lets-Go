package session

import (
	"context"
)

// Define a custom type for keys to avoid potential key collision in context
type stringKey string

// Declare keys for user ID, session ID, and admin access
var (
	userID    stringKey = "userID"
	admin     stringKey = "admin"
	sessionID stringKey = "sessionID"
)

// SetUserID stores the user ID in the context
func SetUserID(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

// SetSessionID stores the session ID in the context
func SetSessionID(ctx context.Context, sID string) context.Context {
	return context.WithValue(ctx, sessionID, sID)
}

// SetAdminAccess stores whether the user has admin access in the context
func SetAdminAccess(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

// GetUserID retrieves the user ID from the context
func GetUserID(ctx context.Context) int {
	// Retrieve the value associated with the user ID key
	if v := ctx.Value(userID); v != nil {
		// Type assert the value to an integer and return it
		if i, ok := v.(int); ok {
			return i
		}
	}
	// Return 0 if user ID is not set or invalid
	return 0
}

// GetSessionID retrieves the session ID from the context
func GetSessionID(ctx context.Context) string {
	// Retrieve the value associated with the session ID key
	if v := ctx.Value(sessionID); v != nil {
		// Type assert the value to a string and return it
		if i, ok := v.(string); ok {
			return i
		}
	}
	return "0"
}

// GetAdmin retrieves the admin access status from the context
func GetAdmin(ctx context.Context) bool {
	if v := ctx.Value(admin); v != nil {
		if i, ok := v.(bool); ok {
			return i
		}
	}
	// Return false if admin access is not set or invalid
	return false
}
