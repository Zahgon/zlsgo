package limiter

import (
	"github.com/sohaha/zlsgo/zsync"
)

type circleQueue struct {
	slice   []int64
	maxSize int
	head    int
	tail    int
	mu      *zsync.RBMutex
}

// newCircleQueue Initialize ring queue
func newCircleQueue(size int) *circleQueue { _ = "STUB: not implemented"; return nil }

func (c *circleQueue) push(val int64) (err error) { _ = "STUB: not implemented"; return nil }

func (c *circleQueue) pop() (val int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *circleQueue) isFull() bool { _ = "STUB: not implemented"; return false }

func (c *circleQueue) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (c *circleQueue) usedSize() int { _ = "STUB: not implemented"; return 0 }

func (c *circleQueue) unUsedSize() int { _ = "STUB: not implemented"; return 0 }

func (c *circleQueue) size() int { _ = "STUB: not implemented"; return 0 }

func (c *circleQueue) deleteExpired() { _ = "STUB: not implemented"; return }
