package zhttp

import (
	"time"
)

func DisableChunked(enable ...bool) { _ = "STUB: not implemented"; return }

func Get(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Post(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Put(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Head(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Options(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Delete(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Patch(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Connect(url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Trace(url string, v ...interface{}) (*Res, error) { _ = "STUB: not implemented"; return nil, nil }

func Do(method, rawurl string, v ...interface{}) (resp *Res, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DoRetry(attempt int, sleep time.Duration, fn func() (*Res, error)) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
