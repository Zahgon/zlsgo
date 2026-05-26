// Package session provides HTTP session management for znet web applications.
// It supports configurable session storage backends and automatic session handling.
package session

import (
	"time"

	"github.com/sohaha/zlsgo/znet"
)

type (
	// Config holds the configuration for session management.
	// It allows customization of cookie settings and session behavior.
	Config struct {
		CookieName string
		ExpiresAt  time.Duration
		AutoRenew  bool
	}
)

// Default creates a new session handler with the default memory store.
// It accepts optional configuration functions to customize the session behavior.
// The returned handler can be used as middleware in znet applications.
func Default(opt ...func(*Config)) znet.Handler {
	_ = "STUB: not implemented"
	return *new(znet.Handler)
}

// New creates a new session handler with the specified store implementation.
// It allows customizing session behavior through configuration options.
// The handler manages session lifecycle including creation, retrieval, and renewal.
func New(stores Store, opt ...func(*Config)) znet.Handler {
	_ = "STUB: not implemented"
	return *new(znet.Handler)
}

// Get retrieves the current session from the context.
// It returns the session if it exists, or an error if no session is found.
// This function is typically used within request handlers to access session data.
func Get(c *znet.Context) (s Session, err error) {
	_ = "STUB: not implemented"
	return *new(Session), nil
}
