package zsync

// NewCoalescer returns a function that coalesces multiple calls into a single
// or a few executions: while one execution is running, further calls schedule
// at least one more run. Thread-safe.
func NewCoalescer(fn func()) func() { _ = "STUB: not implemented"; return nil }
