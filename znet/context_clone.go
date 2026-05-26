package znet

import (
	"net/http"
	"net/url"
)

// Clone creates a request-scoped copy of the context for isolated handler execution.
// Mutable fields are deep-copied so the clone can be used safely in another goroutine.
func (c *Context) Clone(w http.ResponseWriter, req *http.Request) *Context {
	_ = "STUB: not implemented"
	return nil
}

func cloneValues(values url.Values) url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

// CopyResponse copies the prepared response state from another context.
func (c *Context) CopyResponse(from *Context) { _ = "STUB: not implemented"; return }
