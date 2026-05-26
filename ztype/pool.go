package ztype

import (
	"sync"
)

// mapSlicePool []map[string]interface{} slice pool for ToMaps
var mapSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]map[string]interface{}, 0, 4)
	},
}

// stringSlicePool string slice pool for tag option parsing
var stringSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0, 8)
	},
}

// interfaceSlicePool interface{} slice pool for type conversions
var interfaceSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]interface{}, 0, 8)
	},
}

// intSlicePool int slice pool for numeric conversions
var intSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 8)
	},
}

// getMapSlice gets a map slice from object pool
func getMapSlice() []map[string]interface{} { _ = "STUB: not implemented"; return nil }

// putMapSlice returns a map slice to object pool
func putMapSlice(s []map[string]interface{}) { _ = "STUB: not implemented"; return }

// getStringSlice gets a string slice from object pool
func getStringSlice() []string { _ = "STUB: not implemented"; return nil }

// putStringSlice returns a string slice to object pool
func putStringSlice(s []string) { _ = "STUB: not implemented"; return }

// getInterfaceSlice gets an interface{} slice from object pool
func getInterfaceSlice() []interface{} { _ = "STUB: not implemented"; return nil }

// putInterfaceSlice returns an interface{} slice to object pool
func putInterfaceSlice(s []interface{}) { _ = "STUB: not implemented"; return }

// getIntSlice gets an int slice from object pool
func getIntSlice() []int { _ = "STUB: not implemented"; return nil }

// putIntSlice returns an int slice to object pool
func putIntSlice(s []int) { _ = "STUB: not implemented"; return }
