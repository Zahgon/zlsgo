//go:build windows
// +build windows

package zlog

import (
	"syscall"
)

var (
	winEnable          bool
	procSetConsoleMode *syscall.LazyProc
)

func init() {
	if supportColor || isMsystem {
		return
	}
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")

	winEnable = tryApplyOnCONOUT()
	if !winEnable {
		winEnable = tryApplyStdout()
	}
}

func tryApplyOnCONOUT() bool { _ = "STUB: not implemented"; return false }

func tryApplyStdout() bool { _ = "STUB: not implemented"; return false }

func EnableTerminalProcessing(stream syscall.Handle, enable bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IsSupportColor IsSupportColor
func IsSupportColor() bool { _ = "STUB: not implemented"; return false }
