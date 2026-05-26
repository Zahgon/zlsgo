//go:build !amd64 && !arm64 && !ppc64 && !ppc64le && !mips64 && !mips64le && !s390x && !riscv64 && !loong64

package zsync

import "sync"

// RBMutex fallback implementation for non-64-bit architectures.
// It preserves the API but uses a plain RWMutex without read-bias optimizations.
type RBMutex struct {
	rw sync.RWMutex
}

type RBToken struct {
	p *uint64
}

func NewRBMutex() *RBMutex { _ = "STUB: not implemented"; return nil }

func (mu *RBMutex) RLock() RBToken { _ = "STUB: not implemented"; return *new(RBToken) }

func (mu *RBMutex) RUnlock(_ RBToken) { _ = "STUB: not implemented"; return }

func (mu *RBMutex) Lock()   { _ = "STUB: not implemented"; return }
func (mu *RBMutex) Unlock() { _ = "STUB: not implemented"; return }
