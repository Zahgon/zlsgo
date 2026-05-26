//go:build !go1.18
// +build !go1.18

package zshell

import (
	"context"
	"io"
)

func CallbackRunContext(ctx context.Context, command string, callback func(str string, isStdout bool), opt ...func(o *Options)) (<-chan int, func(string), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CallbackRun(command string, callback func(out string, isBasic bool), opt ...func(o *Options)) (<-chan int, func(string), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func Run(command string, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func RunContext(ctx context.Context, command string, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func BgRun(command string, opt ...func(o *Options)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func BgRunContext(ctx context.Context, command string, opt ...func(o *Options)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func OutRun(command string, stdIn io.Reader, stdOut io.Writer, stdErr io.Writer, opt ...func(o Options) Options) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}
