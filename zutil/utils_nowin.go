//go:build !windows
// +build !windows

package zutil

const (
	darwinOpenMax = 10240
)

func IsDoubleClickStartUp() bool { _ = "STUB: not implemented"; return false }

func GetParentProcessName() (string, error) {
	_ = "STUB: not implemented"

	// MaxRlimit tries to set the resource limit RLIMIT_NOFILE to the max (hard limit)
	return "", nil
}

func MaxRlimit() (int, error) { _ = "STUB: not implemented"; return 0, nil }
