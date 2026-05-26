//go:build go1.18
// +build go1.18

package zsync

import (
	"context"
)

func PromiseAllContext[T any](ctx context.Context, promises ...*Promise[T]) *Promise[[]T] {
	_ = "STUB: not implemented"
	return nil
}

func PromiseAll[T any](promises ...*Promise[T]) *Promise[[]T] {
	_ = "STUB: not implemented"
	return nil
}

func PromiseRaceContext[T any](ctx context.Context, promises ...*Promise[T]) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func PromiseRace[T any](promises ...*Promise[T]) *Promise[T] { _ = "STUB: not implemented"; return nil }

type AggregateError struct {
	Errors []error
}

func (ae *AggregateError) Error() string { _ = "STUB: not implemented"; return "" }

func PromiseAnyContext[T any](ctx context.Context, promises ...*Promise[T]) *Promise[T] {
	_ = "STUB: not implemented"
	return nil
}

func PromiseAny[T any](promises ...*Promise[T]) *Promise[T] { _ = "STUB: not implemented"; return nil }
