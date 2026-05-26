//go:build go1.18
// +build go1.18

package zsync

import (
	"unsafe"
)

// SeqLockT is a typed sequence lock that avoids interface conversions on the hot path.
// It provides the same semantics as the untyped SeqLock but returns/accepts T directly.
type SeqLock[T any] struct {
	seq   uint64
	ptr   unsafe.Pointer
	_pad  [56]byte
	_pad2 [56]byte
}

// NewSeqLock creates a typed sequence lock.
func NewSeqLock[T any]() *SeqLock[T] { _ = "STUB: not implemented"; return nil }

// Write publishes a new value with seqlock semantics.
func (s *SeqLock[T]) Write(v T) { _ = "STUB: not implemented"; return }

// Read returns a consistent snapshot if the sequence was stable.
// It may spin briefly under write contention.
func (s *SeqLock[T]) Read() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// writer active
