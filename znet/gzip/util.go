package gzip

import (
	"compress/gzip"
)

type (
	poolCap struct {
		c chan *gzip.Writer
		l int
	}
	// Config gzip configuration
	Config struct {
		// CompressionLevel gzip compression level to use
		CompressionLevel int
		// PoolMaxSize maximum number of resource pools
		PoolMaxSize int
		// MinContentLength minimum content length to trigger gzip, the unit is in byte.
		MinContentLength int
	}
)

func (bp *poolCap) Get() (g *gzip.Writer, err error) { _ = "STUB: not implemented"; return nil, nil }

func (bp *poolCap) Put(g *gzip.Writer) { _ = "STUB: not implemented"; return }
