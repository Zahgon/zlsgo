package zlog

import (
	"fmt"
	"go/ast"
	"io"
	"reflect"
	"strings"
	"sync"
)

type indentWriter struct {
	w   io.Writer
	pre [][]byte
	sel int
	off int
	bol bool
}

func newIndentWriter(w io.Writer, pre ...[]byte) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (w *indentWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func argName(arg ast.Expr) string { _ = "STUB: not implemented"; return "" }

func argNames(filename string, line int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var stringBuilderPool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

func getStringBuilder() *strings.Builder { _ = "STUB: not implemented"; return nil }

func putStringBuilder(sb *strings.Builder) { _ = "STUB: not implemented"; return }

func exprToString(arg ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (fo formatter) String() string { _ = "STUB: not implemented"; return "" }

func (fo formatter) Format(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (fo formatter) passThrough(f fmt.State, c rune) { _ = "STUB: not implemented"; return }

func (p *zprinter) indent() *zprinter { _ = "STUB: not implemented"; return nil }

func (p *zprinter) printInline(v reflect.Value, x interface{}, showType bool) {
	_ = "STUB: not implemented"
	return
}

func (p *zprinter) printStruct(v reflect.Value, showType bool) (stop bool) {
	_ = "STUB: not implemented"
	return false
}

func (p *zprinter) printValue(v reflect.Value, showType, quote bool) {
	_ = "STUB: not implemented"
	return
}

func (p *zprinter) fmtString(s string, quote bool) { _ = "STUB: not implemented"; return }
