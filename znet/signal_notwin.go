//go:build !windows
// +build !windows

package znet

import (
	"github.com/sohaha/zlsgo/zutil"
)

var isRestarting = zutil.NewBool(false)

// Restart triggers a server process restart
func (e *Engine) Restart() error { _ = "STUB: not implemented"; return nil }

// IsRestarting returns whether the server is currently restarting
func (e *Engine) IsRestarting() bool { _ = "STUB: not implemented"; return false }
