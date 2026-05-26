//go:build !windows
// +build !windows

package daemon

import (
	"os"
)

func KillSignal() bool { _ = "STUB: not implemented"; return false }

func SignalChan() (<-chan os.Signal, func()) { _ = "STUB: not implemented"; return nil, nil }

func IsSudo() bool { _ = "STUB: not implemented"; return false }
