//go:build windows
// +build windows

package zutil

func IsDoubleClickStartUp() bool { _ = "STUB: not implemented"; return false }

func GetParentProcessName() (string, error) { _ = "STUB: not implemented"; return "", nil }

// MaxRlimit (not relevant on Windows)
func MaxRlimit() (int, error) { _ = "STUB: not implemented"; return 0, nil }
