package zcache

import (
	"time"

	"github.com/sohaha/zlsgo/ztype"
)

// simple is a global cache instance for convenient access to caching functionality
var simple = NewFast()

// Set adds or updates an item in the global cache with the specified key, value, and optional expiration.
// If no expiration is provided, the item will not expire.
func Set(key string, val interface{}, expiration ...time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Delete removes an item with the specified key from the global cache.
func Delete(key string) {
	_ = "STUB: not implemented"

	// Get retrieves an item from the global cache by its key.
	// Returns the item's value and a boolean indicating whether the item was found.
	return
}

func Get(key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil,

		// GetAny retrieves an item from the global cache and wraps it in a ztype.Type for type conversion.
		// Returns the wrapped value and a boolean indicating whether the item was found.
		false
}

func GetAny(key string) (ztype.Type, bool) {
	_ = "STUB: not implemented"
	return *new(ztype.Type), false
}

// ProvideGet retrieves an item from the global cache, or computes and stores it if not present.
// If the item doesn't exist, the provide function is called to generate the value.
// Returns the item's value and a boolean indicating whether the item was found or created.
func ProvideGet(key string, provide func() (interface{}, bool), expiration ...time.Duration) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
