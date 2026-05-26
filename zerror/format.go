package zerror

import (
	"bytes"
	"fmt"

	"github.com/sohaha/zlsgo/zutil"
)

// Error returns msg
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap returns err inside
func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

// Format formats the frame according to the fmt.Formatter interface
func (e *Error) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Stack returns the stack callers as string
func (e *Error) Stack() string { _ = "STUB: not implemented"; return "" }

// formatSubStack formats the stack for error
func formatSubStack(st zutil.Stack, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }
