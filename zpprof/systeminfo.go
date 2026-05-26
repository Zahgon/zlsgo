package zpprof

import (
	"time"
)

// SystemInfo SystemInfo
type SystemInfo struct {
	ServerName   string
	Runtime      string // runtime duration
	GoroutineNum string // goroutine count
	CPUNum       string // cpu core count
	UsedMem      string // current memory usage
	TotalMem     string // total allocated memory
	SysMem       string // system memory usage
	Lookups      string // pointer lookup count
	Mallocs      string // memory allocation count
	Frees        string // memory release count
	LastGCTime   string // time since last GC
	NextGC       string // next GC memory reclaim amount
	PauseTotalNs string // total GC pause time
	PauseNs      string // last GC pause time
	HeapInuse    string // heap memory in use
}

func NewSystemInfo(startTime time.Time) *SystemInfo { _ = "STUB: not implemented"; return nil }
