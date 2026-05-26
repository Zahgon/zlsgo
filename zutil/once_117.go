//go:build !go1.18
// +build !go1.18

package zutil

// Once initialize the singleton
func Once(fn func() interface{}) func() interface{} { _ = "STUB: not implemented"; return nil }
