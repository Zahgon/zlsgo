// Package zpool provides a thread-safe work pool implementation.
// It allows for concurrent execution of tasks with configurable limits.
package zpool

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sohaha/zlsgo/zdi"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	// Task Define function callbacks
	Task     interface{}
	taskfn   func() error
	WorkPool struct {
		workers     sync.Pool
		injector    zdi.Injector
		queue       *workerQueue
		closeCh     chan struct{}
		usedNum     *zutil.Int64
		activeNum   *zutil.Int64
		panicFunc   PanicFunc
		New         func()
		minIdle     uint
		maxIdle     uint
		releaseTime time.Duration
		mu          sync.RWMutex
		closed      bool
	}
	worker struct {
		jobQueue  chan taskfn
		stop      chan struct{}
		Parameter chan []interface{}
		state     int32
		queueElem *list.Element
	}
	PanicFunc func(err error)
)

var (
	ErrPoolClosed  = errors.New("pool has been closed")
	ErrWaitTimeout = errors.New("pool wait timeout")
)

const (
	workerStateIdle int32 = iota
	workerStateBusy
	workerStateClosing
)

type workerQueue struct {
	mu    sync.Mutex
	list  *list.List
	ready chan struct{}
}

func newWorkerQueue() *workerQueue { _ = "STUB: not implemented"; return nil }

func (q *workerQueue) push(w *worker) { _ = "STUB: not implemented"; return }

func (q *workerQueue) tryPop() *worker { _ = "STUB: not implemented"; return nil }

func (q *workerQueue) pop(ctx context.Context, stop <-chan struct{}) (*worker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *workerQueue) remove(w *worker) bool { _ = "STUB: not implemented"; return false }

// type Options func(*WorkPool)
// // func WithReleaseTime
// func NewCustom(min int, opt Options) *WorkPool {
// 	w := New(min)
// 	if opt != nil {
// 		opt(w)
// 	}
// 	return w
// }

func New(size int, max ...int) *WorkPool { _ = "STUB: not implemented"; return nil }

// TODO: periodically write queue to chan

// Do Add to the workpool and implement
func (wp *WorkPool) Do(fn Task) error { _ = "STUB: not implemented"; return nil }

func (wp *WorkPool) DoWithTimeout(fn Task, t time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// PanicFunc Do Add to the workpool and implement
func (wp *WorkPool) PanicFunc(handler PanicFunc) { _ = "STUB: not implemented"; return }

func (wp *WorkPool) do(cxt context.Context, fn taskfn, param []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// IsClosed Has it been closed
func (wp *WorkPool) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Close  the pool
func (wp *WorkPool) Close() { _ = "STUB: not implemented"; return }

// Wait for the task to finish
func (wp *WorkPool) Wait() { _ = "STUB: not implemented"; return }

// Pause pause
func (wp *WorkPool) Pause() {
	_ = "STUB: not implemented"

	// Continue to work
	return
}

func (wp *WorkPool) Continue(workerNum ...int) { _ = "STUB: not implemented"; return }

// Cap get the number of coroutines
func (wp *WorkPool) Cap() uint { _ = "STUB: not implemented"; return 0 }

// AdjustSize adjust the pool size
func (wp *WorkPool) AdjustSize(workSize int) { _ = "STUB: not implemented"; return }

func (wp *WorkPool) PreInit() error { _ = "STUB: not implemented"; return nil }

func (w *worker) createGoroutines(wp *WorkPool, q *workerQueue, handler PanicFunc) {
	_ = "STUB: not implemented"
	return
}

// case parameter := <-w.Parameter:
// 	q <- w

func (w *worker) close() { _ = "STUB: not implemented"; return }
