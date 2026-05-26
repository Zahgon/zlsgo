package zlsgo

import (
	"testing"
)

const unableCallerInfo = "<Unable to get information>"

// testReporter defines the minimal reporting methods used by TestUtil.
type testReporter interface {
	Helper()
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Fatal(args ...interface{})
}

// TestUtil wraps testing.TB and provides lightweight assertions, error checks,
// logging helpers, and subtest helpers.
//
// Use NewTest to create a helper in unit tests:
//
//	tt := zlsgo.NewTest(t)
//	tt.Equal(1, 1)
//	tt.NoError(err)
//	tt.Len(3, []int{1, 2, 3})
//	tt.Run("sub", func(tt *zlsgo.TestUtil) {
//		tt.EqualTrue(true)
//	})
//
// Assertions cover equality, nil and error checks, string containment,
// length checks, and panic checks.
// Passing true in exit makes an assertion fail the current test immediately.
// Methods T, Run, and Parallel require the underlying value to be *testing.T.
type TestUtil struct {
	tb       testing.TB
	reporter testReporter
}

// NewTest creates a TestUtil from testing.TB.
//
// The returned helper uses the testing logger for assertion output and is
// intended for test assertions plus lightweight subtest orchestration.
// Methods T, Run, and Parallel require the underlying value to be *testing.T.
func NewTest(t testing.TB) *TestUtil { _ = "STUB: not implemented"; return nil }

// newTestUtil builds a TestUtil with a custom reporter.
func newTestUtil(tb testing.TB, reporter testReporter) *TestUtil {
	_ = "STUB: not implemented"
	return nil
}

// GetCallerInfo returns the file name and line number of the test caller.
func (u *TestUtil) GetCallerInfo() string { _ = "STUB: not implemented"; return "" }

// Equal compares expected and actual values with reflect.DeepEqual.
// Passing true in exit will stop the current test immediately on failure.
func (u *TestUtil) Equal(expected, actual interface{}, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// NoEqual compares expected and actual values and asserts they are not equal.
func (u *TestUtil) NoEqual(expected, actual interface{}, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// EqualTrue asserts that the actual value is true.
func (u *TestUtil) EqualTrue(actual interface{}, exit ...bool) { _ = "STUB: not implemented"; return }

// EqualFalse asserts that the actual value is false.
func (u *TestUtil) EqualFalse(actual interface{}, exit ...bool) { _ = "STUB: not implemented"; return }

// True asserts that the actual boolean value is true.
func (u *TestUtil) True(actual bool, exit ...bool) bool { _ = "STUB: not implemented"; return false }

// False asserts that the actual boolean value is false.
func (u *TestUtil) False(actual bool, exit ...bool) bool { _ = "STUB: not implemented"; return false }

// EqualNil asserts that the actual value is nil.
func (u *TestUtil) EqualNil(actual interface{}, exit ...bool) { _ = "STUB: not implemented"; return }

// NoError asserts that err is nil and reports failures through testing output.
func (u *TestUtil) NoError(err error, exit ...bool) bool { _ = "STUB: not implemented"; return false }

// Error asserts that err is not nil.
func (u *TestUtil) Error(err error, exit ...bool) bool { _ = "STUB: not implemented"; return false }

// ErrorContains asserts that err is not nil and its message contains expected.
func (u *TestUtil) ErrorContains(expected string, err error, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// EqualExit compares expected and actual values and immediately fails the test if not equal.
func (u *TestUtil) EqualExit(expected, actual interface{}) { _ = "STUB: not implemented"; return }

// Contains asserts that actual contains expected.
func (u *TestUtil) Contains(expected, actual string, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// NotContains asserts that actual does not contain expected.
func (u *TestUtil) NotContains(expected, actual string, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Len asserts that the actual value length matches the expected length.
func (u *TestUtil) Len(expected int, actual interface{}, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Panics asserts that fn panics.
func (u *TestUtil) Panics(fn func(), exit ...bool) bool { _ = "STUB: not implemented"; return false }

// NotPanics asserts that fn does not panic.
func (u *TestUtil) NotPanics(fn func(), exit ...bool) bool { _ = "STUB: not implemented"; return false }

// Log logs the given values to the test output.
func (u *TestUtil) Log(v ...interface{}) { _ = "STUB: not implemented"; return }

// Logf logs the formatted string to the test output.
func (u *TestUtil) Logf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Fatal logs the given values to the test output and immediately fails the test.
func (u *TestUtil) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

// PrintMyName returns the caller information for the current test.
func (u *TestUtil) PrintMyName() string { _ = "STUB: not implemented"; return "" }

// Run runs a subtest with the given name and function.
// It is only available when the underlying testing object is *testing.T.
func (u *TestUtil) Run(name string, f func(tt *TestUtil)) { _ = "STUB: not implemented"; return }

// T returns the underlying *testing.T object.
// It fails immediately if TestUtil was not created from *testing.T.
func (u *TestUtil) T() *testing.T { _ = "STUB: not implemented"; return nil }

// IsNil asserts that the actual value is nil.
func (u *TestUtil) IsNil(actual interface{}, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// NotNil asserts that the actual value is not nil.
func (u *TestUtil) NotNil(actual interface{}, exit ...bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Parallel marks the test as a parallel test.
// It is only available when the underlying testing object is *testing.T.
func (u *TestUtil) Parallel() { _ = "STUB: not implemented"; return }

// TestCase represents a test case with a name and arbitrary data.
type TestCase struct {
	Data interface{}
	Name string
}

// RunTests runs a series of named test cases through subtests.
func (u *TestUtil) RunTests(tests []TestCase, testFunc func(tt *TestUtil, tc TestCase)) {
	_ = "STUB: not implemented"
	return
}

// ErrorTestCase represents a test case with error expectations.
type ErrorTestCase struct {
	Input    interface{}
	Expected interface{}
	Name     string
	WantErr  bool
}

// RunErrorTests runs test cases for functions returning (result, error).
func (u *TestUtil) RunErrorTests(
	tests []ErrorTestCase,
	testFunc func(input interface{}) (interface{}, error),
) {
	_ = "STUB: not implemented"
	return
}

// failAssertion reports an assertion failure and respects exit behavior.
func (u *TestUtil) failAssertion(exit []bool, format string, args ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// requireTestingT returns the underlying *testing.T when available.
func (u *TestUtil) requireTestingT(method string) *testing.T { _ = "STUB: not implemented"; return nil }

// shouldExit reports whether the assertion should stop the test.
func shouldExit(exit []bool) bool { _ = "STUB: not implemented"; return false }

// valueType returns the reflected type name of v.
func valueType(v interface{}) string { _ = "STUB: not implemented"; return "" }

// valuesEqual compares two values with reflect.DeepEqual.
func valuesEqual(expected, actual interface{}) bool { _ = "STUB: not implemented"; return false }

// isNilValue reports whether v is nil or a typed nil value.
func isNilValue(v interface{}) bool { _ = "STUB: not implemented"; return false }

// lengthOf returns the length of supported collection values.
func lengthOf(v interface{}) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// catchPanic runs fn and captures any panic value.
func catchPanic(fn func()) (panicked bool, recovered interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}
