package zlog

import (
	"bytes"
	"go/ast"
	"go/token"
	"io"
	"sync"
	"time"
)

type astCacheEntry struct {
	File    *ast.File
	ModTime time.Time
	FileSet *token.FileSet
}

var astCache = struct {
	entries map[string]astCacheEntry
	sync.RWMutex
}{
	entries: make(map[string]astCacheEntry),
}

var fileSetPool = sync.Pool{
	New: func() interface{} {
		return token.NewFileSet()
	},
}

func getFileSet() *token.FileSet { _ = "STUB: not implemented"; return nil }

func putFileSet(fset *token.FileSet) { _ = "STUB: not implemented"; return }

func parseFileWithCache(filename string) (*ast.File, *token.FileSet, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var zprinterPool = sync.Pool{
	New: func() interface{} {
		p := new(zprinter)
		p.visited = make(map[visit]int)
		return p
	},
}

func getZprinter(w io.Writer) *zprinter { _ = "STUB: not implemented"; return nil }

func putZprinter(p *zprinter) { _ = "STUB: not implemented"; return }

func formatDateAppend(buf *bytes.Buffer, t time.Time) { _ = "STUB: not implemented"; return }

func formatTimeAppend(buf *bytes.Buffer, t time.Time) { _ = "STUB: not implemented"; return }
