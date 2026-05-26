//go:build !go1.18
// +build !go1.18

package zstring

import (
	"sync"
)

// rngPool is a pool of random number generators to reduce allocation overhead
var rngPool sync.Pool

// Uint32 generates a pseudorandom uint32 value using a simple xorshift algorithm.
// It initializes the generator state from the current time if needed.
func (r *ru) Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

// RandUint32 returns a pseudorandom uint32 value.
// It uses a pool of generators to improve performance by reducing allocations.
func RandUint32() uint32 { _ = "STUB: not implemented"; return 0 }
