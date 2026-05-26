package znet

import (
	"context"
	"io"

	"github.com/sohaha/zlsgo/zutil"
)

// SSE represents a Server-Sent Events connection.
// It handles the event stream between the server and client.
type SSE struct {
	ctx       context.Context
	events    chan *sseEvent
	net       *Context
	option    *SSEOption
	ctxCancel context.CancelFunc
	flush     func()
	stopping  *zutil.Bool
	lastID    string
	method    string
	Comment   []byte
}

// sseEvent represents a single Server-Sent Event with its components.
type sseEvent struct {
	ID      string // Event identifier
	Event   string // Event type
	Comment string // Event comment
	Data    []byte // Event data payload
}

// LastEventID returns the ID of the last event sent over this SSE connection.
func (s *SSE) LastEventID() string {
	_ = "STUB: not implemented"

	// Done returns a channel that's closed when the SSE connection is terminated.
	// This can be used to detect when the client disconnects.
	return ""
}

func (s *SSE) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Stop terminates the SSE connection.
// This will close the event stream and release associated resources.
func (s *SSE) Stop() { _ = "STUB: not implemented"; return }

// sendComment sends a ping comment to keep the connection alive.
func (s *SSE) sendComment() { _ = "STUB: not implemented"; return }

func (s *SSE) Send(id string, data string, event ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SSE) Push() { _ = "STUB: not implemented"; return }

// Use memory manager's pooled buffer for better performance

// SendByte sends raw byte data as an SSE event.
// It allows specifying an event ID and optional event type.
func (s *SSE) SendByte(id string, data []byte, event ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// SSEOption defines configuration options for an SSE connection.
type SSEOption struct {
	RetryTime      int // Client reconnection time in milliseconds
	HeartbeatsTime int // Heartbeat interval in seconds
}

// NewSSE creates a new Server-Sent Events connection from an HTTP context.
// It configures the connection based on the provided options and starts the event loop.
func NewSSE(c *Context, opts ...func(lastID string, opts *SSEOption)) *SSE {
	_ = "STUB: not implemented"
	return nil
}

// RetryTime:      3000,

// Stream sends a streaming response to the client.
// The provided step function is called repeatedly until it returns false.
// Each call to step should write data to the provided writer.
func (c *Context) Stream(step func(w io.Writer) bool) bool { _ = "STUB: not implemented"; return false }
