// Package zlog provides a flexible logging service with support for different log levels,
// colored output, file logging, and debugging utilities.
package zlog

import (
	"fmt"
	"os"

	"github.com/sohaha/zlsgo/zutil"
)

var (
	log    = NewZLog(os.Stdout, "", BitDefault, LogDump, true, 4)
	osExit = func(code int) {
		if zutil.IsDoubleClickStartUp() {
			_, _ = fmt.Scanln()
		}
		os.Exit(code)
	}
)

// SetDefault sets the default logger instance used by the package-level logging functions.
// This allows customizing the global logger behavior.
func SetDefault(l *Logger) { _ = "STUB: not implemented"; return }

// GetFlags returns the current flag bits controlling the format of log output.
// Deprecated: please use SetDefault and access the logger's methods directly.
func GetFlags() int { _ = "STUB: not implemented"; return 0 }

// DisableConsoleColor turns off colored output in the console.
// Deprecated: please use SetDefault and access the logger's methods directly.
func DisableConsoleColor() { _ = "STUB: not implemented"; return }

// ForceConsoleColor forces colored output in the console even when output is not a terminal.
// Deprecated: please use SetDefault and access the logger's methods directly.
func ForceConsoleColor() { _ = "STUB: not implemented"; return }

// ResetFlags sets the output flags to control the formatting of log messages.
// Deprecated: please use SetDefault and access the logger's methods directly.
func ResetFlags(flag int) { _ = "STUB: not implemented"; return }

// AddFlag adds the specified flag to the current set of output flags.
// Deprecated: please use SetDefault and access the logger's methods directly.
func AddFlag(flag int) {
	_ = "STUB: not implemented"

	// SetPrefix sets the prefix for each log line output.
	// Deprecated: please use SetDefault and access the logger's methods directly.
	return
}

func SetPrefix(prefix string) { _ = "STUB: not implemented"; return }

// SetFile configures logging to a file at the specified path.
// The archive parameter controls whether to archive old logs.
// Deprecated: please use SetDefault and access the logger's methods directly.
func SetFile(filepath string, archive ...bool) { _ = "STUB: not implemented"; return }

// SetSaveFile configures logging to a file at the specified path.
// The archive parameter controls whether to archive old logs.
// Deprecated: please use SetDefault and access the logger's methods directly.
func SetSaveFile(filepath string, archive ...bool) { _ = "STUB: not implemented"; return }

func SetLevelFile(level int, filepath string, archive ...bool) { _ = "STUB: not implemented"; return }

func SetLevelSaveFile(level int, filepath string, archive ...bool) {
	_ = "STUB: not implemented"
	return
}

// SetLogLevel sets the minimum level of messages that will be logged.
// Deprecated: please use SetDefault and access the logger's methods directly.
func SetLogLevel(level int) { _ = "STUB: not implemented"; return }

// GetLogLevel returns the current minimum log level.
// Deprecated: please use SetDefault and access the logger's methods directly.
func GetLogLevel() int {
	_ = "STUB: not implemented"

	// Debugf logs a formatted debug message.
	return 0
}

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Debug logs a debug message.
func Debug(v ...interface{}) {
	_ = "STUB: not implemented"

	// Dump logs detailed information about variables, including their names when possible.
	// This is useful for debugging complex data structures.
	return
}

func Dump(v ...interface{}) { _ = "STUB: not implemented"; return }

// Successf logs a formatted success message.
func Successf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Success logs a success message.
func Success(v ...interface{}) {
	_ = "STUB: not implemented"

	// Infof logs a formatted informational message.
	return
}

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Info logs an informational message.
func Info(v ...interface{}) {
	_ = "STUB: not implemented"

	// Tipsf logs a formatted tip message.
	return
}

func Tipsf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Tips logs a tip message.
func Tips(v ...interface{}) {
	_ = "STUB: not implemented"

	// Warnf logs a formatted warning message.
	return
}

func Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Warn logs a warning message.
func Warn(v ...interface{}) {
	_ = "STUB: not implemented"

	// Errorf logs a formatted error message.
	return
}

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Error logs an error message.
func Error(v ...interface{}) {
	_ = "STUB: not implemented"

	// Printf logs a formatted message with no specific level.
	return
}

func Printf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Println logs a message with no specific level.
func Println(v ...interface{}) {
	_ = "STUB: not implemented"

	// Fatalf logs a formatted fatal error message and terminates the program.
	return
}

func Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Fatal logs a fatal error message and terminates the program.
func Fatal(v ...interface{}) {
	_ = "STUB: not implemented"

	// Panicf logs a formatted error message and panics.
	return
}

func Panicf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

// Panic logs an error message and panics.
func Panic(v ...interface{}) {
	_ = "STUB: not implemented"

	// Track logs the current function call stack with the given message.
	// The optional integer parameter controls how many levels of the stack to skip.
	return
}

func Track(v string, i ...int) {
	_ = "STUB: not implemented"

	// Stack logs the full stack trace along with the given value.
	// This is useful for debugging complex call paths and understanding execution flow.
	return
}

func Stack(v interface{}) {
	_ = "STUB: not implemented"

	// Discard sets the default logger to discard all output.
	// This is useful for silencing logs in tests or when logs are not needed.
	return
}

func Discard() { _ = "STUB: not implemented"; return }
