//go:build go1.18
// +build go1.18

package zshell

import (
	"context"
	"io"
)

func toCommand[T string | []string](command T) []string { _ = "STUB: not implemented"; return nil }

func CallbackRunContext[T string | []string](ctx context.Context, command T, callback func(str string, isStdout bool), opt ...func(o *Options)) (<-chan int, func(string), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CallbackRun[T string | []string](command T, callback func(out string, isBasic bool), opt ...func(o *Options)) (<-chan int, func(string), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func Run[T string | []string](command T, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func RunContext[T string | []string](ctx context.Context, command T, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func BgRun[T string | []string](command T, opt ...func(o *Options)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func BgRunContext[T string | []string](ctx context.Context, command T, opt ...func(o *Options)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func OutRun[T string | []string](command T, stdIn io.Reader, stdOut io.Writer, stdErr io.Writer, opt ...func(o Options) Options) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}
