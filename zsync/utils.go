package zsync

import (
	_ "unsafe"
)

const (
	cacheLineSize = 64
)

// nextPowOf2 returns the next power of 2 greater than or equal to v.
// This is used internally for sizing data structures that perform better
// with power-of-2 sizes, such as hash tables and lock arrays.
func nextPowOf2(v uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// parallelism returns the effective parallelism level for the current process.
// It returns the minimum of GOMAXPROCS and the number of CPU cores,
// which provides a reasonable estimate of the available concurrent execution capacity.
func parallelism() uint32 { _ = "STUB: not implemented"; return 0 }
