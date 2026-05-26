package timeout

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"time"

	"github.com/sohaha/zlsgo/znet"
)

type bufferedResponseWriter struct {
	body        bytes.Buffer
	base        http.ResponseWriter
	header      http.Header
	code        int
	wroteHeader bool
}

func newBufferedResponseWriter(base http.ResponseWriter) *bufferedResponseWriter {
	_ = "STUB: not implemented"
	return nil
}

func (w *bufferedResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *bufferedResponseWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *bufferedResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (w *bufferedResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func (w *bufferedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func New(waitingTime time.Duration, custom ...znet.HandlerFunc) znet.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(znet.HandlerFunc)
}

func applyBufferedResponse(target, child *znet.Context, writer *bufferedResponseWriter) {
	_ = "STUB: not implemented"
	return
}
