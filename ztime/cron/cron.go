package cron

import (
	"sync"
	"time"
)

type (
	// Job represents a scheduled task
	// Contains the cron expression, next execution time, and the function to execute
	Job struct {
		expr     *Expression // Parsed cron expression
		NextTime time.Time   // Next time the job will run
		run      func()      // Function to execute
		mu       sync.Mutex  // Mutex to protect NextTime field during concurrent access
	}

	// JobTable manages multiple scheduled tasks
	// Provides functionality to add, run, and stop tasks
	JobTable struct {
		table        sync.Map // Thread-safe map for storing tasks
		sync.RWMutex          // Read-write mutex to protect the stop field
		stop         bool     // Flag indicating whether the job table has been stopped
	}
)

// New creates and returns a new JobTable instance
// Used for managing scheduled tasks
func New() *JobTable {
	_ = "STUB: not implemented"

	// Add adds a new scheduled task to the job table
	// cronLine parameter is a standard cron expression, e.g., "0 * * * * *" means run every minute
	// fn parameter is the function to execute
	// Returns a function to remove the task and a possible error
	return nil
}

func (c *JobTable) Add(cronLine string, fn func()) (remove func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ForceRun immediately checks and executes all due tasks.
func (c *JobTable) ForceRun() (nextTime time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Run starts the task scheduler, beginning to execute tasks according to their schedule.
func (c *JobTable) Run(block ...bool) { _ = "STUB: not implemented"; return }

// Stop stops the task scheduler.
func (c *JobTable) Stop() { _ = "STUB: not implemented"; return }
