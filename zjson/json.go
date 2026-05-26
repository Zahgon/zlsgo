// Package zjson provides fast and flexible JSON manipulation functions.
// It offers path-based operations for getting, setting, and modifying JSON data
// without the need for intermediate unmarshaling and marshaling.
package zjson

import (
	"errors"
	"unsafe"
)

// Error definitions for common JSON operations
var (
	// ErrNoChange is returned when an operation doesn't modify the JSON
	ErrNoChange = errors.New("no change")
	// ErrPathEmpty is returned when an empty path is provided
	ErrPathEmpty = errors.New("path cannot be empty")
	// ErrInvalidJSON is returned when the input is not valid JSON
	ErrInvalidJSON = errors.New("invalid json")
	// ErrNotAllowedWildcard is returned when a wildcard is used in a path where not allowed
	ErrNotAllowedWildcard = errors.New("wildcard characters not allowed in path")
	// ErrNotAllowedArrayAccess is returned when array access is used in a path where not allowed
	ErrNotAllowedArrayAccess = errors.New("array access character not allowed in path")
	// ErrTypeError is returned when the JSON value is not of the expected type
	ErrTypeError = errors.New("json must be an object or array")
)

// MatchKeys returns a new Res containing only the key-value pairs where the key
// matches one of the provided keys.
func (r *Res) MatchKeys(keys []string) *Res { _ = "STUB: not implemented"; return nil }

// Filter returns a new Res containing only the key-value pairs that satisfy
// the provided filter function.
func (r *Res) Filter(fn func(key, value *Res) bool) *Res { _ = "STUB: not implemented"; return nil }

// stringHeader represents the header of a string for unsafe pointer operations.
type stringHeader struct {
	data unsafe.Pointer
	len  int
}

// fillIndex calculates the index of the value within the original JSON string.
func fillIndex(json string, c *parseContext) { _ = "STUB: not implemented"; return }

// set performs the core JSON modification operation, handling various types of modifications.
// It supports setting values, deleting values, and optimistic path resolution.
func set(s, path, raw string, stringify, del, optimistic, place bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use pooled buffer for better performance

// Copy result to avoid returning pooled buffer

// SetOptions sets a JSON value at the specified path with custom options.
// It returns the modified JSON string and any error encountered.
func SetOptions(json, path string, value interface{},
	opts *Options,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetBytesOptions sets a JSON value at the specified path with custom options.
// It works directly with byte slices for better performance.
func SetBytesOptions(json []byte, path string, value interface{},
	opts *Options,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetRawBytesOptions sets a raw JSON value at the specified path in a JSON byte slice with custom options.
// It accepts raw JSON bytes for both the target JSON and the value to be set.
func SetRawBytesOptions(json []byte, path string, value []byte,
	opts *Options,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func safeInt(f float64) (n int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func squash(json string) string { _ = "STUB: not implemented"; return "" }

func parseSquash(json string, i int) (string, int) { _ = "STUB: not implemented"; return "", 0 }

// switchJson processes a JSON string starting at position i, handling nested structures.
// It returns the processed JSON string and the new position.
func switchJson(json string, i int, isParse bool) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}
