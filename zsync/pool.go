//go:build go1.18
// +build go1.18

package zsync

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func NewPool[T any](n func() T) *Pool[T] { _ = "STUB: not implemented"; return nil }

func (p *Pool[T]) Get() T { _ = "STUB: not implemented"; return *new(T) }

func (p *Pool[T]) Put(x T) { _ = "STUB: not implemented"; return }
