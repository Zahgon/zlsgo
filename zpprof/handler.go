package zpprof

import (
	"time"

	"github.com/sohaha/zlsgo/znet"
)

var startTime = time.Now()

func infoHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func indexHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func allocsHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func mutexHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func heapHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func goroutineHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func blockHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func threadCreateHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func cmdlineHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func profileHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func symbolHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func traceHandler(c *znet.Context) { _ = "STUB: not implemented"; return }

func redirectPprof(c *znet.Context) { _ = "STUB: not implemented"; return }

func authDebug(token string) znet.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(znet.HandlerFunc)
}
