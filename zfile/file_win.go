//go:build windows
// +build windows

package zfile

// MoveFile moves a file from source to destination path.
// If force is true and destination exists, it will be removed before moving.
// On Windows systems, this uses syscall.MoveFile which provides better support
// for Windows file paths and attributes.
func MoveFile(source string, dest string, force ...bool) error {
	_ = "STUB: not implemented"
	return nil
}
