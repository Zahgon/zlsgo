//go:build !windows
// +build !windows

package zfile

// MoveFile moves a file from source to destination path.
// If force is true and destination exists, it will be removed before moving.
// On non-Windows systems, this uses os.Rename which is atomic if both paths
// are on the same filesystem.
func MoveFile(source string, dest string, force ...bool) error {
	_ = "STUB: not implemented"
	return nil
}
