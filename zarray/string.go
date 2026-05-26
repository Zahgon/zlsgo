//go:build go1.18
// +build go1.18

package zarray

// Slice splits a string by the specified separator and converts each part to type T.
// Empty parts after trimming whitespace are excluded from the result.
// If n is provided, the string will be split into at most n parts.
// Returns an empty slice if the input string is empty.
func Slice[T comparable](s, sep string, n ...int) []T { _ = "STUB: not implemented"; return nil }

// Join concatenates the elements of a slice into a single string with the specified separator.
// Empty string elements are excluded from the result.
// Returns an empty string if the input slice is empty.
func Join[T comparable](s []T, sep string) string { _ = "STUB: not implemented"; return "" }
