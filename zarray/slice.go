//go:build go1.18
// +build go1.18

package zarray

// CopySlice creates and returns a new slice containing all elements from the input slice.
func CopySlice[T any](l []T) []T { _ = "STUB: not implemented"; return nil }

// Rand returns a random element from the provided slice.
// If the slice is empty, returns the zero value of type T.
func Rand[T any](collection []T) T { _ = "STUB: not implemented"; return *new(T) }

// RandPickN returns a new slice containing n randomly selected elements from the input collection.
// If n is greater than the collection length, returns all elements in random order.
// If n is less than or equal to 0 or the collection is empty, returns an empty slice.
func RandPickN[T any](collection []T, n int) []T { _ = "STUB: not implemented"; return nil }

// Map applies the iteratee function to each element in the collection and returns a new slice
// containing the transformed values. If parallel is provided, the operation is performed
// concurrently using the specified number of workers.
func Map[T any, R any](collection []T, iteratee func(int, T) R, parallel ...uint) []R {
	_ = "STUB: not implemented"
	return nil
}

// ParallelMap applies the iteratee function to each element in the collection concurrently
// and returns a new slice containing the transformed values.
// If the calculation does not involve time-consuming operations, using Map is recommended.
// Deprecated: please use Map with a parallel parameter instead
func ParallelMap[T any, R any](collection []T, iteratee func(int, T) R, workers uint) []R {
	_ = "STUB: not implemented"
	return nil
}

// Shuffle creates and returns a new slice containing all elements from the input slice
// in a random order. The original slice remains unchanged.
func Shuffle[T any](collection []T) []T { _ = "STUB: not implemented"; return nil }

// Reverse creates and returns a new slice containing all elements from the input slice
// in reverse order. The original slice remains unchanged.
func Reverse[T any](collection []T) []T { _ = "STUB: not implemented"; return nil }

// Filter creates a new slice containing all elements from the input slice that satisfy
// the predicate function. The original slice remains unchanged.
// Filter creates a new slice containing only the elements that satisfy the predicate.
// This optimized version pre-allocates capacity based on the input slice size,
// avoiding unnecessary full copy when only a subset of elements are kept.
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	_ = "STUB: not implemented"
	// Pre-allocate with a reasonable capacity estimate
	// Using 1/4 of original size as initial capacity to reduce allocations
	// while avoiding excessive memory usage for highly selective filters
	return nil
}

// Contains checks if a value exists in the collection.
// Returns true if the value is found, false otherwise.
func Contains[T comparable](collection []T, v T) bool { _ = "STUB: not implemented"; return false }

// Find searches for an element in the slice that satisfies the predicate function.
// Returns the found element and true if successful, or the zero value and false if not found.
func Find[T any](collection []T, predicate func(index int, item T) bool) (res T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Unique creates and returns a new slice containing only the unique elements
// from the input slice, preserving the original order of first occurrence.
func Unique[T comparable](collection []T) []T { _ = "STUB: not implemented"; return nil }

// Diff compares two slices and returns two new slices containing the elements that
// are unique to each input slice (not present in the other slice).
func Diff[T comparable](list1 []T, list2 []T) ([]T, []T) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pop removes and returns the last element from the slice.
// If the slice is empty, returns the zero value of type T.
// This function modifies the original slice.
func Pop[T comparable](list *[]T) (v T) { _ = "STUB: not implemented"; return *new(T) }

// Shift removes and returns the first element from the slice.
// If the slice is empty, returns the zero value of type T.
// This function modifies the original slice.
func Shift[T comparable](list *[]T) (v T) { _ = "STUB: not implemented"; return *new(T) }

// Chunk splits the slice into multiple sub-slices of the specified size.
// The last chunk may contain fewer elements if the slice length is not divisible by size.
// If size is less than or equal to 0 or the slice is empty, returns an empty slice of slices.
func Chunk[T any](slice []T, size int) [][]T { _ = "STUB: not implemented"; return nil }

// RandShift returns a closure function that, when called, returns a random element from the list.
// Each element is returned exactly once in random order. When all elements have been returned,
// subsequent calls will return an error. The original list is not modified.
func RandShift[T comparable](list []T) func() (T, error) { _ = "STUB: not implemented"; return nil }

// SortWithPriority sorts the slice based on the priority of elements.
// The elements in the 'first' slice are placed at the beginning of the result,
// followed by the elements in the 'last' slice, and then the remaining elements.
func SortWithPriority[T comparable](slice []T, first, last []T) []T {
	_ = "STUB: not implemented"
	return nil
}
