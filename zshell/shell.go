// Package zshell use a simple way to execute shell commands
package zshell

import (
	"bytes"
	"context"
	"io"
	"os/exec"
)

var (
	Debug = false
	Env   []string
	Dir   string
)

type ShellBuffer struct {
	writer io.Writer
	buf    *bytes.Buffer
}

func newShellStdBuffer(writer io.Writer) *ShellBuffer { _ = "STUB: not implemented"; return nil }

func (s *ShellBuffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *ShellBuffer) String() string { _ = "STUB: not implemented"; return "" }

func ExecCommandHandle(ctx context.Context, command []string,
	bef func(cmd *exec.Cmd) error, aft func(cmd *exec.Cmd, err error)) (code int,
	err error,
) {
	_ = "STUB: not implemented"
	return 0, nil
}

func cmdResult(cmd *exec.Cmd) (code int, isSuccess bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type pipeWork struct {
	cmd *exec.Cmd
	// r   *io.PipeReader
	w *io.PipeWriter
}

func PipeExecCommand(ctx context.Context, commands [][]string, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func ExecCommand(ctx context.Context, command []string, stdIn io.Reader, stdOut io.Writer, stdErr io.Writer, opt ...func(o *Options)) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

type Options struct {
	Dir        string
	Env        []string
	CloseStdin bool
}

func callbackRunContext(ctx context.Context, commandArgs []string, callback func(str string, isStdout bool), opt ...func(o *Options)) (<-chan int, func(string), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func fixCommand(command string) (runCommand []string) { _ = "STUB: not implemented"; return nil }
