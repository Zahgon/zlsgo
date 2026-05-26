package zhttp

import (
	"context"
	"io"
	"net/http"
)

// NewRequest create new request
func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

// GetRequest get request from pool
func (e *Engine) GetRequest() *Request { _ = "STUB: not implemented"; return nil }

// GetRequest get request from pool
func GetRequest() *Request { _ = "STUB: not implemented"; return nil }

// URL set request url
func (r *Request) URL(url string) *Request { _ = "STUB: not implemented"; return nil }

// Method set request method
func (r *Request) Method(method string) *Request { _ = "STUB: not implemented"; return nil }

// Header set request header
func (r *Request) Header(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// Headers set request headers
func (r *Request) Headers(headers Header) *Request { _ = "STUB: not implemented"; return nil }

// Query set query param
func (r *Request) Query(key string, value interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// QueryMap set query params
func (r *Request) QueryMap(params QueryParam) *Request { _ = "STUB: not implemented"; return nil }

// Form set form param
func (r *Request) Form(key string, value interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// FormMap set form params
func (r *Request) FormMap(params Param) *Request { _ = "STUB: not implemented"; return nil }

// Body set request body
func (r *Request) Body(body interface{}) *Request { _ = "STUB: not implemented"; return nil }

// JSON set json request body
func (r *Request) JSON(v interface{}) *Request { _ = "STUB: not implemented"; return nil }

// XML set xml request body
func (r *Request) XML(v interface{}) *Request { _ = "STUB: not implemented"; return nil }

// File set file upload
func (r *Request) File(fieldName, fileName string, file io.ReadCloser) *Request {
	_ = "STUB: not implemented"
	return nil
}

// Client set http client
func (r *Request) Client(client *http.Client) *Request { _ = "STUB: not implemented"; return nil }

// Cookie set cookie
func (r *Request) Cookie(cookie *http.Cookie) *Request { _ = "STUB: not implemented"; return nil }

// Context set context
func (r *Request) Context(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

// Host set host
func (r *Request) Host(host string) *Request { _ = "STUB: not implemented"; return nil }

// UploadProgress set upload progress callback
func (r *Request) UploadProgress(progress UploadProgress) *Request {
	_ = "STUB: not implemented"
	return nil
}

// DownloadProgress set download progress callback
func (r *Request) DownloadProgress(progress DownloadProgress) *Request {
	_ = "STUB: not implemented"
	return nil
}

// NoRedirect disable redirect
func (r *Request) NoRedirect(disable bool) *Request { _ = "STUB: not implemented"; return nil }

// Custom custom request handler
func (r *Request) Custom(fn CustomReq) *Request { _ = "STUB: not implemented"; return nil }

// Do do request
func (r *Request) Do() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// GET do get request
func (r *Request) GET() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// POST do post request
func (r *Request) POST() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// PUT do put request
func (r *Request) PUT() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// PATCH do patch request
func (r *Request) PATCH() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// DELETE do delete request
func (r *Request) DELETE() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// HEAD do head request
func (r *Request) HEAD() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// OPTIONS do options request
func (r *Request) OPTIONS() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// Reset reset request
func (r *Request) Reset() *Request { _ = "STUB: not implemented"; return nil }

// Release release request to pool
func (r *Request) Release() { _ = "STUB: not implemented"; return }

// DoAndRelease do request and release to pool
func (r *Request) DoAndRelease() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// GetAndRelease do get request and release to pool
func (r *Request) GetAndRelease() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

// PostAndRelease do post request and release to pool
func (r *Request) PostAndRelease() (*Res, error) { _ = "STUB: not implemented"; return nil, nil }
