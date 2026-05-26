package zhttp

import (
	"context"
	"net/http"
)

type (
	SSEEngine struct {
		ctx          context.Context
		eventCh      chan *SSEEvent
		errCh        chan error
		ctxCancel    context.CancelFunc
		verifyHeader func(http.Header) bool
		option       SSEOption
		readyState   int
	}

	SSEEvent struct {
		ID        string
		Event     string
		Undefined []byte
		Data      []byte
	}
)

var (
	delim   = []byte{':'} // []byte{':', ' '}
	ping    = []byte("ping")
	dataEnd = byte('\n')
)

func (sse *SSEEngine) Event() <-chan *SSEEvent { _ = "STUB: not implemented"; return nil }

func (sse *SSEEngine) Close() { _ = "STUB: not implemented"; return }

func (sse *SSEEngine) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (sse *SSEEngine) Error() <-chan error { _ = "STUB: not implemented"; return nil }

func (sse *SSEEngine) VerifyHeader(fn func(http.Header) bool) { _ = "STUB: not implemented"; return }

func (sse *SSEEngine) OnMessage(fn func(*SSEEvent)) (<-chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SSE(url string, v ...interface{}) (*SSEEngine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) sseReq(method, url string, v ...interface{}) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SSEOption struct {
	Method   string
	RetryNum int
}

func (e *Engine) SSE(url string, opt func(*SSEOption), v ...interface{}) (*SSEEngine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Continue to retry
