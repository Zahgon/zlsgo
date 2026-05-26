package cache

import (
	"time"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/znet"
)

type (
	// Config configuration
	Config struct {
		Custom func(c *znet.Context) (key string, expiration time.Duration)
		zcache.Options
	}
	cacheContext struct {
		Type    string
		Content []byte
		Code    int32
	}
)

func New(opt ...func(conf *Config)) znet.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(znet.HandlerFunc)
}

func QueryKey(c *znet.Context) (key string) { _ = "STUB: not implemented"; return "" }
