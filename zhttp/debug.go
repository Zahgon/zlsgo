package zhttp

import (
	"bytes"
	"io"
	"net"
	"time"

	"github.com/sohaha/zlsgo/zutil"
)

const (
	rn = "\r\n\r\n"
)

var (
	Debug = zutil.NewBool(false)
)

type dumpConn struct {
	io.Writer
	io.Reader
}

func (c *dumpConn) Close() error                       { _ = "STUB: not implemented"; return nil }
func (c *dumpConn) LocalAddr() net.Addr                { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *dumpConn) RemoteAddr() net.Addr               { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *dumpConn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (c *dumpConn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (c *dumpConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

type (
	dummyBody struct {
		N   int
		off int
	}

	delegateReader struct {
		c chan io.Reader
		r io.Reader
	}
)

func (r *delegateReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *dummyBody) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (d *dummyBody) Close() error { _ = "STUB: not implemented"; return nil }

type dumpBuffer struct {
	bytes.Buffer
}

func (b *dumpBuffer) Write(p []byte) { _ = "STUB: not implemented"; return }

func (b *dumpBuffer) WriteString(s string) { _ = "STUB: not implemented"; return }

func (r *Res) dumpRequest(dump *dumpBuffer) { _ = "STUB: not implemented"; return }

func (r *Res) dumpReqHead(dump *dumpBuffer) { _ = "STUB: not implemented"; return }

func (r *Res) dumpResonse(dump *dumpBuffer) { _ = "STUB: not implemented"; return }

func (r *Res) Cost() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (r *Res) Dump() string { _ = "STUB: not implemented"; return "" }
