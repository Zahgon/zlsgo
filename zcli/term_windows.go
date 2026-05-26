//go:build windows

package zcli

import (
	"os"
	"syscall"
)

type coord struct {
	X int16
	Y int16
}

type smallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        uint16
	Window            smallRect
	MaximumWindowSize coord
}

var (
	kernel32                       = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

func fileIsTerminal(file *os.File) bool { _ = "STUB: not implemented"; return false }

func fileTerminalWidth(file *os.File) (int, bool) { _ = "STUB: not implemented"; return 0, false }
