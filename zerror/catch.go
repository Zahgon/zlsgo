package zerror

// TryCatch exception capture
func TryCatch(fn func() error) (err error) { _ = "STUB: not implemented"; return nil }

// Panic if error is not nil, usually used in conjunction with TryCatch
func Panic(err error) { _ = "STUB: not implemented"; return }
