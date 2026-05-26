// Package zerror provides error related operations
package zerror

import (
	"github.com/sohaha/zlsgo/zutil"
)

type (
	// ErrCode error code type
	ErrCode int32
	// Error wraps err with code
	Error struct {
		err     error
		wrapErr error
		errText *string
		stack   zutil.Stack
		code    ErrCode
		inner   bool
	}

	External func(err error) error
)

var goROOT = zutil.GOROOT()

func New(code ErrCode, text string, w ...External) error { _ = "STUB: not implemented"; return nil }

// Reuse the error
func Reuse(err error) error { _ = "STUB: not implemented"; return nil }

// Wrap wraps err with code
func Wrap(err error, code ErrCode, text string, w ...External) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: please use zerror.With
// SupText returns the error text
func SupText(err error, text string) error { _ = "STUB: not implemented"; return nil }

// With returns the inner error's text
func With(err error, text string, w ...External) error { _ = "STUB: not implemented"; return nil }

// Unwrap returns if err is Error and its code == code
func Unwrap(err error, code ErrCode) (error, bool) { _ = "STUB: not implemented"; return nil, false }

// Is returns if err is Error and its code == code
func Is(err error, code ...ErrCode) bool { _ = "STUB: not implemented"; return false }

// UnwrapCode Returns the current error code
func UnwrapCode(err error) (ErrCode, bool) { _ = "STUB: not implemented"; return *new(ErrCode), false }

// UnwrapCodes Returns the current all error code
func UnwrapCodes(err error) (codes []ErrCode) { _ = "STUB: not implemented"; return nil }

// UnwrapErrors Returns the current all error text
func UnwrapErrors(err error) (errs []string) { _ = "STUB: not implemented"; return nil }

func UnwrapFirst(err error) (ferr error) { _ = "STUB: not implemented"; return nil }

func UnwrapFirstCode(err error) (code ErrCode) { _ = "STUB: not implemented"; return *new(ErrCode) }
