package znet

import (
	"html/template"

	"github.com/sohaha/zlsgo/zlog"
)

// routeLog formats route information for logging purposes.
// It colorizes HTTP methods and pads strings to ensure consistent log formatting.
func routeLog(log *zlog.Logger, tf, method, path string) string {
	_ = "STUB: not implemented"
	return ""
}

// templatesDebug logs information about loaded HTML templates when in debug mode.
// It lists all named templates that have been loaded into the engine.
func templatesDebug(e *Engine, t *template.Template) { _ = "STUB: not implemented"; return }

// routeAddLog logs information about a newly added route when in debug mode.
// It includes the HTTP method, path, handler function name, and middleware count.
func routeAddLog(e *Engine, method string, path string, action Handler, middlewareCount int) {
	_ = "STUB: not implemented"
	return
}
