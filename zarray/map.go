//go:build go1.18
// +build go1.18

package zarray

// Keys extracts all keys from a map and returns them as a slice.
// The order of the keys in the resulting slice is not guaranteed.
func Keys[K comparable, V any](in map[K]V) []K { _ = "STUB: not implemented"; return nil }

// Values extracts all values from a map and returns them as a slice.
// The order of the values in the resulting slice is not guaranteed.
func Values[K comparable, V any](in map[K]V) []V { _ = "STUB: not implemented"; return nil }

// IndexMap indexes a slice of Maps into a map based on a key function.
func IndexMap[K comparable, V any](arr []V, toKey func(V) (K, V)) (map[K]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FlatMap flattens a map of Maps into a single slice of Maps.
func FlatMap[K comparable, V any](m map[K]V, fn func(key K, value V) V) []V {
	_ = "STUB: not implemented"
	return nil
}
