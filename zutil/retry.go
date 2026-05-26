package zutil

import (
	"time"
)

// RetryConf holds configuration options for the retry mechanism.
// It controls the number of retries, intervals between attempts, timeout,
// and backoff strategy.
type RetryConf struct {
	// maxRetry is the maximum number of retries
	maxRetry int
	// Interval is the base interval between retry attempts
	Interval time.Duration
	// MaxRetryInterval is the maximum interval between retry attempts
	// when using exponential backoff
	MaxRetryInterval time.Duration
	// Timeout is the maximum total duration for all retry attempts
	Timeout time.Duration
	// BackOffDelay determines whether to use exponential backoff
	// for increasing the interval between retries
	BackOffDelay bool
}

// DoRetry executes a function with retry logic based on the provided configuration.
// It will retry the function up to 'sum' times or until it succeeds.
// Additional options can be provided to customize retry behavior.
func DoRetry(sum int, fn func() error, opt ...func(*RetryConf)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// BackOffDelay calculates the delay duration for exponential backoff retry strategy.
// It increases the delay exponentially based on the attempt number and adds jitter
// to prevent synchronized retries in distributed systems.
func BackOffDelay(attempt int, retryInterval, maxRetryInterval time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
