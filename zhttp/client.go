package zhttp

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func newClient() *http.Client { _ = "STUB: not implemented"; return nil }

func (e *Engine) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (e *Engine) SetClient(client *http.Client) { _ = "STUB: not implemented"; return }

func (e *Engine) DisableChunked(enable ...bool) { _ = "STUB: not implemented"; return }

func (e *Engine) Get(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Post(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Put(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Patch(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Delete(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Head(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Options(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Trace(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Connect(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) DoRetry(attempt int, sleep time.Duration, fn func() (*Res, error)) (res *Res, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) EnableInsecureTLS(enable bool) { _ = "STUB: not implemented"; return }

// DisableHTTP2 disables HTTP/2 protocol support.
// This can be used to mitigate CVE-2023-45288 (HTTP/2 CONTINUATION flood vulnerability)
// when upgrading golang.org/x/net is not possible.
func (e *Engine) DisableHTTP2() { _ = "STUB: not implemented"; return }

type Certificate struct {
	CertFile string
	KeyFile  string
}

func (e *Engine) TlsCertificate(certs ...Certificate) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) EnableCookie(enable bool) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) CheckRedirect(fn ...func(req *http.Request, via []*http.Request) error) {
	_ = "STUB: not implemented"
	return
}

func (e *Engine) SetTimeout(d time.Duration) { _ = "STUB: not implemented"; return }

func (e *Engine) SetTransport(transport func(*http.Transport)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) SetProxyUrl(proxyUrl ...string) error { _ = "STUB: not implemented"; return nil }

func (e *Engine) SetProxy(proxy func(*http.Request) (*url.URL, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) RemoveProxy() error { _ = "STUB: not implemented"; return nil }

func (e *Engine) getJSONEncOpts() *jsonEncOpts { _ = "STUB: not implemented"; return nil }

func (e *Engine) SetJSONEscapeHTML(escape bool) { _ = "STUB: not implemented"; return }

func (e *Engine) SetJSONIndent(prefix, indent string) { _ = "STUB: not implemented"; return }

func (e *Engine) SetXMLIndent(prefix, indent string) { _ = "STUB: not implemented"; return }

func (e *Engine) SetSsl(certPath, keyPath, CAPath string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) getTransport() *http.Transport { _ = "STUB: not implemented"; return nil }

func (e *Engine) getXMLEncOpts() *xmlEncOpts { _ = "STUB: not implemented"; return nil }

// OptimizeForHighConcurrency High concurrency optimization
func (e *Engine) OptimizeForHighConcurrency() { _ = "STUB: not implemented"; return }

// OptimizeForLowLatency Low latency optimization
func (e *Engine) OptimizeForLowLatency() { _ = "STUB: not implemented"; return }
