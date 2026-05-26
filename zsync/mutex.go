//go:build amd64 || arm64 || ppc64 || ppc64le || mips64 || mips64le || s390x || riscv64 || loong64

package zsync

import (
	"sync"
	_ "unsafe" // required for //go:linkname
)

type RBMutex struct {
	rslots         []optRSlot
	state          uint64
	rw             sync.RWMutex
	rmask          uint32
	writerMomentum uint32
	_              [64]byte
	_              [56]byte
	_              [4]byte
}

type optRSlot struct {
	counter uint64
	_       [56]byte
}

type RBToken struct {
	p *uint64
}

const (
	rbiasShift           = 63
	rbiasMask            = uint64(1) << rbiasShift
	writerShift          = 32
	writerMask           = uint64(0x7FFFFFFF) << writerShift
	defaultBiasLimit     = 4
	writerMomentumMedium = 6
	writerMomentumHigh   = 12
)

// Use runtime's proc pin/unpin to obtain a stable per-P identifier.
// These are internal runtime functions accessed via linkname.
// See: src/runtime/proc.go (procPin/procUnpin)
//
//go:linkname runtime_procPin runtime.procPin
func runtime_procPin() int

//go:linkname runtime_procUnpin runtime.procUnpin
func runtime_procUnpin()

//go:nosplit
func getProcID() uint32 {
	_ = "STUB: not implemented"
	// Pin to current P to retrieve its id, then unpin immediately.
	// We only need the id for slot selection; we do not keep the P pinned
	// across the critical section to avoid excessive overhead.
	return 0
}

//go:nosplit
func likely(b bool) bool {
	_ = "STUB: not implemented"

	//go:nosplit
	return false
}

func unlikely(b bool) bool {
	_ = "STUB: not implemented"

	// NewRBMutex Extreme optimized version of read bias lock (read more and write less scene)
	return false
}

func NewRBMutex() *RBMutex { _ = "STUB: not implemented"; return nil }

//go:nosplit
func (mu *RBMutex) RLock() RBToken { _ = "STUB: not implemented"; return *new(RBToken) }

// If slots are not initialized, fall back to RWMutex to keep semantics safe.

//go:nosplit
func (mu *RBMutex) RUnlock(token RBToken) { _ = "STUB: not implemented"; return }

func (mu *RBMutex) Lock() { _ = "STUB: not implemented"; return }

// Adaptive backoff to reduce potential writer starvation

// As contention persists, yield more aggressively

func (mu *RBMutex) Unlock() { _ = "STUB: not implemented"; return }

// Only enable read-bias if slots are initialized.

//go:nosplit
//go:inline
func biasLimit(momentum uint32) uint32 { _ = "STUB: not implemented"; return 0 }
