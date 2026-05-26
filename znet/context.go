package znet

import (
	"github.com/sohaha/zlsgo/zdi"
)

// Host returns the current host with scheme (http/https).
// If full is true, it includes the request URL path.
func (c *Context) Host(full ...bool) string { _ = "STUB: not implemented"; return "" }

// CompletionLink ensures a URL is absolute by prepending the current host
// if the provided link is relative.
func (c *Context) CompletionLink(link string) string { _ = "STUB: not implemented"; return "" }

// IsWebsocket determines if the current request is a WebSocket upgrade request
// by checking the Connection and Upgrade headers.
func (c *Context) IsWebsocket() bool { _ = "STUB: not implemented"; return false }

// IsSSE determines if the current request is expecting Server-Sent Events
// by checking if the Accept header contains 'text/event-stream'.
func (c *Context) IsSSE() bool { _ = "STUB: not implemented"; return false }

// IsAjax determines if the current request is an AJAX request
// by checking for the X-Requested-With header with value XMLHttpRequest.
func (c *Context) IsAjax() bool { _ = "STUB: not implemented"; return false }

// GetClientIP returns the client's IP address by checking various headers
// and connection information. It attempts to determine the most accurate
// client IP, even when behind proxies.
func (c *Context) GetClientIP() string { _ = "STUB: not implemented"; return "" }

// GetHeader returns the value of the specified request header.
func (c *Context) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

// SetHeader sets a response header with the given key and value.
// If value is empty, the header will be removed.
func (c *Context) SetHeader(key, value string, only ...bool) { _ = "STUB: not implemented"; return }

// write finalizes the response by writing headers and body data to the response writer.
// It handles content negotiation, status codes, and ensures headers are properly set.
func (c *Context) write() { _ = "STUB: not implemented"; return }

// Next executes all remaining middleware in the chain.
// Returns false if the middleware chain has been stopped with Abort().
func (c *Context) Next() bool { _ = "STUB: not implemented"; return false }

// next is an internal method that executes the next middleware in the chain.
// It's called by Next() and handles the middleware execution flow.
func (c *Context) next() {
	_ = "STUB: not implemented"
	// If already terminated, return directly
	return
}

// Check if there are more middleware

// Get current middleware and advance queue

// Execute middleware (outside lock)

// SetCookie sets an HTTP cookie with the given name and value.
// Optional maxAge parameter specifies the cookie's max age in seconds (0 = session cookie).
// Sets SameSite=Lax by default for basic CSRF protection.
func (c *Context) SetCookie(name, value string, maxAge ...int) { _ = "STUB: not implemented"; return }

// SetSecureCookie sets an HTTP cookie with security flags enabled.
// Optional maxAge parameter specifies the cookie's max age in seconds (0 = session cookie).
// This method sets Secure=true and SameSite=Strict by default for enhanced security.
// Use SetCookie directly for more options like Domain, Path customization, etc.
func (c *Context) SetSecureCookie(name, value string, maxAge ...int) {
	_ = "STUB: not implemented"
	return
}

// GetCookie returns the value of the cookie with the given name.
// Returns an empty string if the cookie doesn't exist.
func (c *Context) GetCookie(name string) string { _ = "STUB: not implemented"; return "" }

// GetReferer returns the Referer header of the request, which contains
// the URL of the page that linked to the current page.
func (c *Context) GetReferer() string { _ = "STUB: not implemented"; return "" }

// GetUserAgent returns the User-Agent header of the request, which identifies
// the client software originating the request.
func (c *Context) GetUserAgent() string { _ = "STUB: not implemented"; return "" }

// ContentType returns or sets the Content-Type header.
// If contentText is provided, it sets the Content-Type header.
// Otherwise, it returns the current Content-Type of the request.
func (c *Context) ContentType(contentText ...string) string { _ = "STUB: not implemented"; return "" }

// WithValue stores a key-value pair in the context for sharing data
// between middleware and handlers. Returns the context for chaining.
func (c *Context) WithValue(key string, value interface{}) *Context {
	_ = "STUB: not implemented"
	return nil
}

// Value retrieves data stored in the context by key.
// It returns the value and a boolean indicating if the key exists.
// If the key doesn't exist and default values are provided, the first default is returned.
func (c *Context) Value(key string, def ...interface{}) (value interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// MustValue retrieves data stored in the context by key, with simplified return.
// If the key doesn't exist and default values are provided, the first default is returned.
// Unlike Value(), this method only returns the value without the existence flag.
func (c *Context) MustValue(key string, def ...interface{}) (value interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// Injector returns the dependency injection container for this context.
// This allows handlers to access shared services and dependencies.
func (c *Context) Injector() zdi.Injector {
	_ = "STUB: not implemented"

	// FileAttachment serves a file as an attachment with the specified filename.
	// This will prompt the browser to download the file rather than display it.
	return *new(zdi.Injector)
}

func (c *Context) FileAttachment(filepath, filename string) { _ = "STUB: not implemented"; return }

// isASCII checks if a string contains only ASCII characters.
// Source: https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
func isASCII(s string) bool { _ = "STUB: not implemented"; return false }
