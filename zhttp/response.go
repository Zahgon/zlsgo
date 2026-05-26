package zhttp

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sohaha/zlsgo/zjson"
)

type Res struct {
	err    error
	r      *Engine
	req    *http.Request
	resp   *http.Response
	client *http.Client
	*multipartHelper
	downloadProgress DownloadProgress
	tmpFile          string
	requesterBody    []byte
	responseBody     []byte
	cost             time.Duration
}

func (r *Res) Request() *http.Request { _ = "STUB: not implemented"; return nil }

func (r *Res) Response() *http.Response { _ = "STUB: not implemented"; return nil }

func (r *Res) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (r *Res) GetCookie() map[string]*http.Cookie { _ = "STUB: not implemented"; return nil }

func (r *Res) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (r *Res) Stream(fn func(line []byte, eof bool) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Res) ToBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func forceUTF8(r *Res, respBody []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *Res) Body() (body io.ReadCloser) { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

func (r *Res) HTML() (doc QueryHTML) { _ = "STUB: not implemented"; return *new(QueryHTML) }

func (r *Res) String() string { _ = "STUB: not implemented"; return "" }

func (r *Res) JSONs() *zjson.Res { _ = "STUB: not implemented"; return nil }

func (r *Res) JSON(key string) *zjson.Res { _ = "STUB: not implemented"; return nil }

func (r *Res) ToString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Res) ToJSON(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Res) ToXML(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Res) ToFile(name string) error { _ = "STUB: not implemented"; return nil }

//noinspection GoUnhandledErrorResult

//noinspection GoUnhandledErrorResult

func (r *Res) download(file *os.File) error { _ = "STUB: not implemented"; return nil }

//noinspection GoUnhandledErrorResult
