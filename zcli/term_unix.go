//go:build !windows

package zcli

import (
	"os"
)

type termSize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func fileIsTerminal(file *os.File) bool { _ = "STUB: not implemented"; return false }

func fileTerminalWidth(file *os.File) (int, bool) { _ = "STUB: not implemented"; return 0, false }
