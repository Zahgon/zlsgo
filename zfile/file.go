// Package zfile provides file and path operations for common development tasks.
// It includes utilities for file manipulation, path resolution, and directory management.
package zfile

import (
	"os"
	"strings"
	"time"
)

var ProjectPath = "./"

func init() {
	// abs, _ := filepath.Abs(".")
	// ProjectPath = RealPath(abs)
	ProjectPath = ProgramPath()
	if strings.Contains(ProjectPath, TmpPath("")) {
		ProjectPath = RootPath()
	}
}

// PathExist checks if a path exists and determines if it's a directory or file.
// Returns:
//   - 1: path exists and is a directory
//   - 2: path exists and is a file
//   - 3: path exists and is a symlink
//   - 0: path does not exist
func PathExist(path string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// DirExist checks if the specified path exists and is a directory.
// Returns true if the path exists and is a directory, false otherwise.
func DirExist(path string) bool { _ = "STUB: not implemented"; return false }

// FileExist checks if the specified path exists and is a file.
// Returns true if the path exists and is a file, false otherwise.
func FileExist(path string) bool { _ = "STUB: not implemented"; return false }

// FileSize returns the formatted size of a file (e.g., "1.5 MB").
// If the file doesn't exist, it returns an empty string.
func FileSize(file string) (size string) { _ = "STUB: not implemented"; return "" }

// FileSizeUint returns the size of a file in bytes as a uint64.
// If the file doesn't exist, it returns 0.
func FileSizeUint(file string) (size uint64) { _ = "STUB: not implemented"; return 0 }

// SizeFormat converts a size in bytes (int) to a human-readable string
// with appropriate units (B, KB, MB, GB, etc.).
func SizeFormat(s int64) string { _ = "STUB: not implemented"; return "" }

func logSize(n, b float64) float64 { _ = "STUB: not implemented"; return 0 }

// RootPath returns the absolute path of the current working directory.
// This is typically the directory from which the program was launched.
func RootPath() string { _ = "STUB: not implemented"; return "" }

// TmpPath returns a path to a temporary directory.
// If a pattern is provided, it creates a temporary directory with that pattern.
// Otherwise, it returns the system's temporary directory.
func TmpPath(pattern ...string) string { _ = "STUB: not implemented"; return "" }

// SafePath returns a path that is guaranteed to be within the specified base directory.
// If no base directory is provided, it uses the project path as the base.
// This helps prevent directory traversal vulnerabilities.
func SafePath(path string, pathRange ...string) string { _ = "STUB: not implemented"; return "" }

// RealPath converts a relative path to an absolute path.
// If addSlash is true, it ensures the path ends with a slash.
// The function normalizes path separators to forward slashes.
func RealPath(path string, addSlash ...bool) (realPath string) {
	_ = "STUB: not implemented"
	return ""
}

// RealPathMkdir converts a path to an absolute path and creates the directory if it doesn't exist.
// If addSlash is true, it ensures the path ends with a slash.
// Note: To ensure the directory can be created successfully, use HasReadWritePermission
// to check permissions before calling this function.
func RealPathMkdir(path string, addSlash ...bool) string { _ = "STUB: not implemented"; return "" }

// IsSubPath checks if subPath is contained within path.
// Both paths are converted to absolute paths before comparison.
func IsSubPath(subPath, path string) bool { _ = "STUB: not implemented"; return false }

// Rmdir removes a directory and all its contents recursively.
// If notIncludeSelf is true, it recreates the directory after deletion,
// effectively removing all contents while keeping the directory itself.
func Rmdir(path string, notIncludeSelf ...bool) (ok bool) { _ = "STUB: not implemented"; return false }

// Remove deletes the specified file or empty directory.
// It returns an error if the path doesn't exist or if a non-empty directory is specified.
func Remove(path string) error { _ = "STUB: not implemented"; return nil }

// CopyFile copies a file from source to destination, preserving file mode.
// It creates the destination file if it doesn't exist and overwrites it if it does.
func CopyFile(source string, dest string) (err error) { _ = "STUB: not implemented"; return nil }

// ExecutablePath returns the absolute path of the current executable file.
// If it cannot determine the executable path, it falls back to using os.Args[0].
func ExecutablePath() string { _ = "STUB: not implemented"; return "" }

// ProgramPath returns the directory containing the current executable.
// If addSlash is true, it ensures the path ends with a slash.
// If the executable path cannot be determined, it falls back to ProjectPath.
func ProgramPath(addSlash ...bool) (path string) { _ = "STUB: not implemented"; return "" }

// pathAddSlash adds a trailing slash to the path if addSlash is true and
// the path doesn't already end with a slash.
func pathAddSlash(path string, addSlash ...bool) string { _ = "STUB: not implemented"; return "" }

// GetMimeType determines the MIME type of a file based on its content and/or filename.
// If content is provided, it uses content-based detection first.
// If that fails or returns a generic type, it falls back to extension-based detection.
func GetMimeType(filename string, content []byte) (ctype string) {
	_ = "STUB: not implemented"
	return ""
}

// HasPermission checks if the specified path has the requested permission mode.
// If the path doesn't exist and noUp is false, it checks the parent directory.
// Returns true if the path has the requested permissions, false otherwise.
func HasPermission(path string, perm os.FileMode, noUp ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// HasReadWritePermission checks if the specified path has read and write permissions (0600).
// Returns true if the path has read and write permissions, false otherwise.
func HasReadWritePermission(path string) bool { _ = "STUB: not implemented"; return false }

// fileInfo is an internal structure used to store file metadata for sorting and cleanup operations.
type fileInfo struct {
	// modTime is the last modification time of the file
	modTime time.Time
	// path is the absolute path to the file
	path string
	// size is the file size in bytes
	size uint64
}

// fileInfos is a slice of fileInfo structures that implements sort.Interface
// to allow sorting files by modification time.
type fileInfos []fileInfo

func (f fileInfos) Len() int           { _ = "STUB: not implemented"; return 0 }
func (f fileInfos) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (f fileInfos) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// DirStatOptions provides configuration options for the StatDir function.
type DirStatOptions struct {
	// MaxSize is the maximum allowed total size in bytes for the directory.
	// Files will be deleted (oldest first) if the total size exceeds this value.
	MaxSize uint64
	// MaxTotal is the maximum allowed number of files in the directory.
	// Files will be deleted (oldest first) if the total count exceeds this value.
	MaxTotal uint64
}

// StatDir calculates the total size and number of files in a directory.
// If options are provided with MaxSize or MaxTotal values, it will delete the oldest
// files to keep the directory within the specified limits.
// Returns the final size in bytes, number of files, and any error encountered.
func StatDir(path string, options ...DirStatOptions) (size, total uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
