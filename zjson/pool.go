package zjson

import (
	"sync"
)

// pathCachePool reuses path parsing results
var pathCachePool = sync.Pool{
	New: func() interface{} {
		return make([]pathResult, 0, 8)
	},
}

func getPathCache() []pathResult { _ = "STUB: not implemented"; return nil }

func putPathCache(cache []pathResult) { _ = "STUB: not implemented"; return }
