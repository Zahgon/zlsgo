package limiter

import (
	"time"

	"github.com/sohaha/zlsgo/znet"
)

// Rule user access control strategy
type Rule struct {
	rules []*singleRule
}

// New Newlimiter
func New(allowed uint64, overflow ...func(c *znet.Context)) znet.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(znet.HandlerFunc)
}

// NewRule Custom limiter rule
func NewRule() *Rule {
	_ = "STUB: not implemented"

	// AddRule increase user access control strategy
	// If less than 1s, please use golang.org/x/time/rate
	return nil
}

func (r *Rule) AddRule(exp time.Duration, allowed int, estimated ...int) {
	_ = "STUB: not implemented"
	return
}

// AllowVisit Is access allowed
func (r *Rule) AllowVisit(keys ...interface{}) bool { _ = "STUB: not implemented"; return false }

// AllowVisitByIP AllowVisit IP
func (r *Rule) AllowVisitByIP(ip string) bool { _ = "STUB: not implemented"; return false }
