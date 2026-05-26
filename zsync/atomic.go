//go:build go1.18
// +build go1.18

package zsync

import (
	"sync/atomic"

	"github.com/sohaha/zlsgo/zutil"
)

// AtomicValue is the generic version of [atomic.Value].
type AtomicValue[T any] struct {
	_ zutil.Nocmp
	v atomic.Value
}

type wrappedValue[T any] struct{ v T }

func NewValue[T any](v T) *AtomicValue[T] { _ = "STUB: not implemented"; return nil }

// Load returns the value set by the most recent Store.
// It returns the zero value for T if the value is empty.
func (v *AtomicValue[T]) Load() T { _ = "STUB: not implemented"; return *new(T) }

// Store sets the value of the Value to x.
func (v *AtomicValue[T]) Store(x T) { _ = "STUB: not implemented"; return }

// Swap stores new into Value and returns the previous value.
// It returns the zero value for T if the value was not set before.
func (v *AtomicValue[T]) Swap(x T) (old T) { _ = "STUB: not implemented"; return *new(T) }

// CAS executes the compare-and-swap operation for the Value.
func (v *AtomicValue[T]) CAS(oldV, newV T) (swapped bool) { _ = "STUB: not implemented"; return false }
