//go:build go1.18
// +build go1.18

package zarray

// SortMaper implements an ordered map that maintains insertion order of keys
// while providing map-like operations for key-value pairs.
type SortMaper[K hashable, V any] struct {
	values *Maper[K, V]
	keys   []K
}

// NewSortMap creates a new SortMaper with the specified initial capacity.
// The SortMaper maintains the insertion order of keys while providing map operations.
func NewSortMap[K hashable, V any](size ...uintptr) *SortMaper[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Set adds or updates a key-value pair in the map.
// If the key is new, it is appended to the ordered keys list.
func (s *SortMaper[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

// Get retrieves a value by its key.
// Returns the value and a boolean indicating whether the key was found.
func (s *SortMaper[K, V]) Get(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *

	// Has checks if a key exists in the map.
	// Returns true if the key exists, false otherwise.
	new(V), false
}

func (s *SortMaper[K, V]) Has(key K) (ok bool) { _ = "STUB: not implemented"; return false }

// Delete removes one or more key-value pairs from the map.
// This removes the keys from both the ordered keys list and the underlying map.
func (s *SortMaper[K, V]) Delete(key ...K) { _ = "STUB: not implemented"; return }

// Len returns the number of key-value pairs in the map.
func (s *SortMaper[K, V]) Len() int {
	_ = "STUB: not implemented"

	// Keys returns all keys in the map in their insertion order.
	return 0
}

func (s *SortMaper[K, V]) Keys() []K {
	_ = "STUB: not implemented"

	// ForEach iterates through all key-value pairs in the map in insertion order.
	// The iteration continues as long as the lambda function returns true.
	return nil
}

func (s *SortMaper[K, V]) ForEach(lambda func(K, V) bool) { _ = "STUB: not implemented"; return }
