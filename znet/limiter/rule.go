package limiter

import (
	"sync"
	"time"
)

type singleRule struct {
	notRecordsIndex   map[int]struct{}
	locker            *sync.Mutex
	usedRecordsIndex  sync.Map
	records           []*circleQueue
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	allowed           int
	estimated         int
}

// newRule Initialize an access control policy
func newRule(defaultExpiration time.Duration, allowed int, estimated ...int) *singleRule {
	_ = "STUB: not implemented"
	return nil
}

func createRule(defaultExpiration, cleanupInterval time.Duration, allowed, userEstimated int) *singleRule {
	_ = "STUB: not implemented"
	return nil
}

// allowVisit Whether access is allowed or not. If access is allowed, an access record is added to the access record
func (r *singleRule) allowVisit(key interface{}) bool { _ = "STUB: not implemented"; return false }

// remainingVisits Remaining visits
func (r *singleRule) remainingVisits(key interface{}) int { _ = "STUB: not implemented"; return 0 }

// add access record
func (r *singleRule) add(key interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// deleteExpired Delete expired data
func (r *singleRule) deleteExpired() { _ = "STUB: not implemented"; return }

// deleteExpiredOnce Delete expired data once in a specific time interval
func (r *singleRule) deleteExpiredOnce() { _ = "STUB: not implemented"; return }

func (r *singleRule) recovery() { _ = "STUB: not implemented"; return }

func (r *singleRule) needRecovery() bool { _ = "STUB: not implemented"; return false }
