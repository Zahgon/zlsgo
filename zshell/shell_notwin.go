//go:build !windows
// +build !windows

package zshell

import (
	"context"
	"os/exec"

	"github.com/sohaha/zlsgo/zutil"
)

var chcp = zutil.Once(func() struct{} {
	return struct{}{}
})

func RunNewProcess(file string, args []string) (pid int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func RunBash(ctx context.Context, command string) (code int, outStr, errStr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func sysProcAttr(cmd *exec.Cmd) *exec.Cmd { _ = "STUB: not implemented"; return nil }
