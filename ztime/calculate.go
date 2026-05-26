//go:build go1.18
// +build go1.18

package ztime

import (
	"time"
)

func Diff[T time.Time | string](t1, t2 T, format ...string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// FindRange parses a slice of time strings and returns the earliest and latest time.
func FindRange[T time.Time | string](times []T, format ...string) (time.Time, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), nil
}

func parseGenericTime[T time.Time | string](t T, format ...string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Sequence generates a sequence of time strings between start and end times based on the given format.
func Sequence[T time.Time | string](start, end T, stepFn func(time.Time) time.Time, format ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
