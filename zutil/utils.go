// Package zutil provides utility functions and types for daily development tasks.
package zutil

import (
	"runtime"
	"time"
)

type (
	// Stack represents a call stack as an array of program counters.
	// It provides methods for formatting and analyzing the call stack.
	Stack []uintptr

	// Nocmp is an uncomparable struct that can be embedded in other structs
	// to make them uncomparable (cannot be compared with == or !=).
	Nocmp [0]func()

	// namedArgs is an internal type used to associate a name with an argument value.
	namedArgs struct {
		arg  interface{}
		name string
	}
)

// Named creates a named argument by associating a name with a value.
// This is useful for functions that accept variadic arguments and need to
// distinguish between different argument types or purposes.
func Named(name string, arg interface{}) interface{} { _ = "STUB: not implemented"; return nil }

const (
	// maxStackDepth is the maximum depth of call stack frames to capture.
	maxStackDepth = 1 << 5 // 32 frames
)

// WithRunContext measures the execution time and memory allocation of a function.
// It returns the duration of execution and the number of bytes allocated during execution.
func WithRunContext(handler func()) (time.Duration, int64) {
	_ = "STUB: not implemented"
	return *new(time.Duration), 0
}

// TryCatch executes a function and captures any panic that occurs, converting it to an error.
// If the function returns an error normally, that error is returned.
// If a panic occurs, it is converted to an error and returned.
func TryCatch(fn func() error) (err error) { _ = "STUB: not implemented"; return nil }

// Try executes a function and captures any panic that occurs, passing it to the catch function.
// If a finally function is provided, it is always executed after the main function,
// regardless of whether a panic occurred.
// Deprecated: please use zerror.TryCatch instead.
func Try(fn func(), catch func(e interface{}), finally ...func()) {
	_ = "STUB: not implemented"
	return
}

// CheckErr checks if an error is not nil and panics if it is.
// If exit is true, it prints the error and exits the program instead of panicking.
// Deprecated: please use zerror.Panic instead.
func CheckErr(err error, exit ...bool) { _ = "STUB: not implemented"; return }

// Callers returns a stack trace as a Stack.
// The optional skip parameter indicates how many stack frames to skip before
// starting to collect the stack trace.
func Callers(skip ...int) Stack { _ = "STUB: not implemented"; return *new(Stack) }

var (
	h = []byte{104, 97}
	s = []byte{115}
	o = []byte{111}
	g = []byte{103}
	u = []byte{116, 104, 117, 98, 46, 99, 111, 109, 47}
	l string
	t string
)

func init() {
	l = string(append([]byte{103, 105}, append(u, append(s, append(o, append(h, h...)...)...)...)...))
	t = "_test." + string(append(g, o...))
}

// Format iterates through the stack frames and calls the provided function for each frame.
// The function receives the runtime.Func object, file name, and line number for each frame.
// If the function returns false, iteration stops.
// Note: Frames from the zlsgo library itself are automatically skipped.
func (s Stack) Format(f func(fn *runtime.Func, file string, line int) bool) {
	_ = "STUB: not implemented"
	return
}
