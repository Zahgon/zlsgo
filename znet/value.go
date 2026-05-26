package znet

import (
	"github.com/sohaha/zlsgo/zvalid"
)

// Content-Type MIME of the most common data formats.
// These constants define the standard MIME types used for HTTP content negotiation.
const (
	// mimeJSON is the MIME type for JSON data
	mimeJSON = "application/json"
	// mimePlain is the MIME type for plain text
	mimePlain = "text/plain"
	// mimePOSTForm is the MIME type for URL-encoded form data
	mimePOSTForm = "application/x-www-form-urlencoded"
	// mimeMultipartPOSTForm is the MIME type for multipart form data (typically used for file uploads)
	mimeMultipartPOSTForm = "multipart/form-data"
)

// valid is an internal helper function that validates struct fields using the provided validation rules.
// It supports validation for basic types like string, bool, numeric types, etc.
func (c *Context) valid(obj interface{}, v map[string]zvalid.Engine) error {
	_ = "STUB: not implemented"
	return nil
}

// BindValid binds request data to the provided object and validates it using the provided validation rules.
// It first binds the data using the appropriate method based on request type, then validates the bound data.
func (c *Context) BindValid(obj interface{}, v map[string]zvalid.Engine) error {
	_ = "STUB: not implemented"
	return nil
}
