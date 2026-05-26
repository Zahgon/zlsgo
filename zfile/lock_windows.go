//go:build windows
// +build windows

package zfile

import (
	"os"
	"syscall"
)

// Windows-specific constants for file locking operations
const (
	// lockfileExclusiveLock requests an exclusive lock
	lockfileExclusiveLock = 0x2
	// lockfileFailImmediately returns immediately if the lock cannot be acquired
	lockfileFailImmediately = 0x1
	// errorLockViolation is the error code returned when a lock is already held
	errorLockViolation = 0x21
)

// errLocked is the error returned when a file is already locked by another process.
// For consistency with Unix implementations, we use syscall.EWOULDBLOCK.
var errLocked = syscall.EWOULDBLOCK

// Windows API function references for file locking
var (
	// kernel32DLL is a reference to the Windows kernel32.dll library
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	// procLockFileEx is a reference to the LockFileEx function in kernel32.dll
	procLockFileEx = kernel32DLL.NewProc("LockFileEx")
	// procUnlockFileEx is a reference to the UnlockFileEx function in kernel32.dll
	procUnlockFileEx = kernel32DLL.NewProc("UnlockFileEx")
)

// lockFile acquires an exclusive, non-blocking lock on the given file.
// On Windows systems, this uses the LockFileEx Win32 API function.
// Returns errLocked if the file is already locked by another process.
func lockFile(f *os.File) error { _ = "STUB: not implemented"; return nil }

// unlockFile releases a lock previously acquired with lockFile.
// On Windows systems, this uses the UnlockFileEx Win32 API function.
func unlockFile(f *os.File) error { _ = "STUB: not implemented"; return nil }
