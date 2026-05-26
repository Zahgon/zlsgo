//go:build go1.18
// +build go1.18

package zutil

// Once creates a function that ensures the provided initialization function
// is executed only once, regardless of how many times the returned function is called.
// This implements the singleton pattern with built-in error recovery.
//
// If the initialization function panics, the Once state is reset after a delay,
// allowing for a retry on the next call.
func Once[T any](fn func() T) func() T { _ = "STUB: not implemented"; return nil }

// OnceWithError creates a function that ensures the provided initialization function
// is executed only once, regardless of how many times the returned function is called.
// This implements the singleton pattern with built-in error recovery.
func OnceWithError[T any](fn func() (T, error)) func() (T, error) {
	_ = "STUB: not implemented"
	return nil
}

// Guard creates a function that ensures mutually exclusive execution of the provided function.
// If the returned function is called while a previous call is still in progress,
// it will return an error instead of executing the function again.
//
// This is useful for preventing concurrent execution of functions that are not thread-safe
// or for rate-limiting access to resources.
func Guard[T any](fn func() T) func() (T, error) { _ = "STUB: not implemented"; return nil }
