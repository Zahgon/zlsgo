package zcli

import (
	"io"
)

func terminalWidth(writer io.Writer) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func isTerminalWriter(writer io.Writer) bool { _ = "STUB: not implemented"; return false }

func envTerminalWidth() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func stringDisplayWidth(s string) int { _ = "STUB: not implemented"; return 0 }

func fitProgressLine(prefix, core, suffix string, termWidth int) string {
	_ = "STUB: not implemented"
	return ""
}

func joinProgressLine(prefix, core, suffix string) string { _ = "STUB: not implemented"; return "" }

func truncateDisplayWidth(s string, max int) string { _ = "STUB: not implemented"; return "" }

func nextDisplayToken(s string) (string, int, int) { _ = "STUB: not implemented"; return "", 0, 0 }

func emojiClusterWidth(s string) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func isKeycapBase(r rune) bool { _ = "STUB: not implemented"; return false }

func isEmojiModifier(r rune) bool { _ = "STUB: not implemented"; return false }

func isEmojiBase(r rune) bool { _ = "STUB: not implemented"; return false }

func runeDisplayWidth(r rune) int { _ = "STUB: not implemented"; return 0 }
