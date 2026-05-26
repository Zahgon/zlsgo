// Package zarray provides comprehensive array operations and utilities for working with
// dynamic arrays in Go. It implements various array manipulation functions including
// insertion, deletion, searching, and transformation operations.
package zarray

import (
	"errors"
)

// Array represents a dynamic array that supports insertion, deletion, and random access
// operations. All elements are stored as interface{} type for maximum flexibility.
type Array struct {
	data []interface{}
	size int
}

// ErrIllegalIndex is returned when an operation is attempted with an invalid array index
var ErrIllegalIndex = errors.New("illegal index")

// NewArray initializes a new Array with the specified capacity.
// If no capacity is provided, a default capacity of 5 is used.
func NewArray(capacity ...int) (array *Array) { _ = "STUB: not implemented"; return nil }

// Deprecated: New is deprecated, please use NewArray instead
func New(capacity ...int) (array *Array) { _ = "STUB: not implemented"; return nil }

// CopyArray creates a new Array by copying all elements from the provided array.
// Returns the new Array and any error that occurred during copying.
func CopyArray(arr interface{}) (array *Array, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: Copy is deprecated, please use CopyArray instead
func Copy(arr interface{}) (array *Array, err error) {
	_ = "STUB: not implemented"
	return nil,

		// checkIndex determines whether the provided index is out of bounds.
		// Returns true if the index is invalid, along with the current size of the array.
		nil
}

func (arr *Array) checkIndex(index int) (bool, int) { _ = "STUB: not implemented"; return false, 0 }

// resize expands the array's capacity to the specified size by creating
// a new underlying array and copying all existing elements.
func (arr *Array) resize(capacity int) { _ = "STUB: not implemented"; return }

// CapLength returns the current capacity of the array
func (arr *Array) CapLength() int { _ = "STUB: not implemented"; return 0 }

// Length returns the current number of elements in the array
func (arr *Array) Length() int {
	_ = "STUB: not implemented"

	// IsEmpty returns true if the array contains no elements, false otherwise
	return 0
}

func (arr *Array) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Unshift inserts an element at the beginning of the array.
// Returns an error if the operation fails.
func (arr *Array) Unshift(value interface{}) error { _ = "STUB: not implemented"; return nil }

// Push appends one or more elements to the end of the array
func (arr *Array) Push(values ...interface{}) { _ = "STUB: not implemented"; return }

// Add inserts an element at the specified index position.
// Returns an error if the index is out of bounds or if the operation fails.
func (arr *Array) Add(index int, value interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If the current number of elements is equal to the arr capacity,
// the arr will be expanded to twice the original size

// Map creates a new array by applying the provided function to each element.
// The function receives the index and value of each element and returns the transformed value.
func (arr *Array) Map(fn func(int, interface{}) interface{}) *Array {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves the element at the specified index position.
// If the index is invalid and a default value is provided, returns the default value.
// Otherwise returns the element and any error that occurred.
func (arr *Array) Get(index int, def ...interface{}) (value interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set modifies the element at the specified index position.
// Returns an error if the index is out of bounds.
func (arr *Array) Set(index int, value interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Contains checks if the specified value exists in the array.
// Returns true if found, false otherwise.
func (arr *Array) Contains(value interface{}) bool { _ = "STUB: not implemented"; return false }

// Index finds the position of the specified value in the array.
// Returns the index (in range [0, n-1]) if found, or -1 if not found.
func (arr *Array) Index(value interface{}) int { _ = "STUB: not implemented"; return 0 }

// Remove deletes one or more elements starting at the specified index position.
// Returns the removed elements and any error that occurred during the operation.
func (arr *Array) Remove(index int, l ...int) (value []interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shift removes and returns the first element of the array.
// Returns the removed element and any error that occurred during the operation.
func (arr *Array) Shift() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// Pop removes and returns the last element of the array.
		// Returns the removed element and any error that occurred during the operation.
		nil
}

func (arr *Array) Pop() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// RemoveValue removes the first occurrence of the specified element from the array.
// Returns the removed element and any error that occurred during the operation.
func (arr *Array) RemoveValue(value interface{}) (e interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clear removes all elements from the array, resetting it to an empty state
func (arr *Array) Clear() { _ = "STUB: not implemented"; return }

// Raw returns a copy of the underlying array data as a slice of interface{} values
func (arr *Array) Raw() []interface{} { _ = "STUB: not implemented"; return nil }

// Format returns a string representation of the array including its size, capacity, and elements
func (arr *Array) Format() (format string) { _ = "STUB: not implemented"; return "" }

// Shuffle creates a new array with the same elements in random order
func (arr *Array) Shuffle() (array *Array) { _ = "STUB: not implemented"; return nil }

// GetInf retrieves the element at the specified index from a slice of interface{} values.
// If the index is invalid and a default value is provided, returns the default value.
// Otherwise returns the element and any error that occurred.
func GetInf(arr []interface{}, index int, def ...interface{}) (value interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
