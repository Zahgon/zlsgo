// Package zpprof provides a register for zweb framework to use net/http/pprof easily.
package zpprof

import (
	"github.com/sohaha/zlsgo/znet"
)

// Register Registration routing
func Register(r *znet.Engine, token string) (RouterGroup *znet.Engine) {
	_ = "STUB: not implemented"

	// go tool pprof http://127.0.0.1:8081/debug/pprof/profile
	// go tool pprof -alloc_space http://127.0.0.1:8081/debug/pprof/heap
	// go tool pprof -inuse_space http://127.0.0.1:8081/debug/pprof/heap
	return nil
}

func ListenAndServe(addr ...string) error { _ = "STUB: not implemented"; return nil }
