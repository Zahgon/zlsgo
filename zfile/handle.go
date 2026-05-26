package zfile

// CopyDir recursively copies the source directory to the destination directory.
// It preserves file permissions and can filter files using the optional filterFn function.
// If the filter function returns false for a file, that file will not be copied.
func CopyDir(source string, dest string, filterFn ...func(srcFilePath, destFilePath string) bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReadFile reads the entire contents of a file into memory.
// It returns the file contents as a byte slice and any error encountered.
func ReadFile(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ReadLineFile reads a file line by line and calls the provided handle function for each line.
// The handle function receives the line number and the line content as parameters.
// If the handle function returns an error, reading stops and that error is returned.
func ReadLineFile(path string, handle func(line int, data []byte) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// WriteFile writes data to a file, creating the file if it doesn't exist.
// If isAppend is true, data is appended to the file; otherwise, the file is overwritten.
// It creates any necessary parent directories automatically.
func WriteFile(path string, b []byte, isAppend ...bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PutOffset writes data to a file at the specified offset position.
// If the file doesn't exist, it will be created.
// This is useful for modifying specific portions of a file without rewriting the entire content.
func PutOffset(path string, b []byte, offset int64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PutAppend appends data to the end of a file.
// If the file doesn't exist, it will be created along with any necessary parent directories.
// This is a convenience wrapper around WriteFile with append mode.
func PutAppend(path string, b []byte) (err error) { _ = "STUB: not implemented"; return nil }
