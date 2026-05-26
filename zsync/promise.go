//go:build go1.18
// +build go1.18

package zsync

import (
	"context"

	"github.com/sohaha/zlsgo/zutil"
)

// PromiseState represents the current state of a Promise.
type PromiseState uint8

// Promise represents an asynchronous operation that may produce a value or an error.
// It implements a Promise pattern similar to JavaScript promises but adapted for Go.
// The type parameter T represents the type of the value that will be produced.
type Promise[T any] struct {
	value   T               // The resolved value
	done    chan struct{}   // Channel that is closed when the promise is settled
	reason  error           // The rejection reason if the promise is rejected
	timeout *zutil.Bool     // Flag indicating if the promise has timed out
	ctx     context.Context // Context for cancellation
}

// Catch registers a function to be called when the Promise is rejected.
// It returns a new Promise that resolves with the result of the callback,
// or rejects with the error returned by the callback.
func (p *Promise[T]) Catch(rejected func(error) (T, error)) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

// Finally registers a function to be called when the Promise is settled,
// regardless of whether it was fulfilled or rejected.
// The function receives no arguments and its return value is ignored.
func (p *Promise[T]) Finally(finally func()) *Promise[T] { _ = "STUB: not implemented"; return nil }

// Then registers callbacks to be called when the Promise is settled.
// If the Promise is fulfilled, the fulfilled callback is called with the value.
// If the Promise is rejected and a rejected callback is provided, it is called with the error.
// It returns a new Promise that resolves with the result of the appropriate callback.
func (p *Promise[T]) Then(fulfilled func(T) (T, error), rejected ...func(error) (T, error)) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

// Done waits for the Promise to be settled and returns the result.
// If the Promise was fulfilled, it returns the value and nil error.
// If the Promise was rejected, it returns the zero value and the error.
// This method blocks until the Promise is settled or the context is canceled.
func (p *Promise[T]) Done() (value T, reason error) { _ = "STUB: not implemented"; return *new(T), nil }

// NewPromiseContext creates a new Promise with the given context and executor function.
// The executor function is called immediately in a new goroutine.
// The Promise will be fulfilled with the value returned by the executor,
// or rejected with the error returned by the executor.
// The Promise will be rejected if the context is canceled.
func NewPromiseContext[T any](ctx context.Context, executor func() (T, error)) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewPromise creates a new Promise with the given executor function.
// It uses context.Background() as the context.
// See NewPromiseContext for more details.
func NewPromise[T any](executor func() (T, error)) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}
