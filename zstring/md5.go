package zstring

// ProjectMd5 calculates the MD5 hash of the current executable.
// This is useful for checking if the application binary has been modified.
func ProjectMd5() string { _ = "STUB: not implemented"; return "" }

// Md5 calculates the MD5 hash of a string.
// It returns the hash as a hexadecimal encoded string.
func Md5(s string) string { _ = "STUB: not implemented"; return "" }

// Md5Byte calculates the MD5 hash of a byte slice.
// It returns the hash as a hexadecimal encoded string.
func Md5Byte(s []byte) string { _ = "STUB: not implemented"; return "" }

// Md5File calculates the MD5 hash of a file at the given path.
// It returns the hash as a hexadecimal encoded string and any error encountered.
func Md5File(path string) (encrypt string, err error) { _ = "STUB: not implemented"; return "", nil }

// r := bufio.NewReader(f)
