package cors

import (
	"net/http"
	"sync"

	"github.com/sohaha/zlsgo/znet"
)

type (
	Config struct {
		CustomHandler Handler
		methods       string
		credentials   string
		headers       string
		exposeHeaders string
		Domains       []string
		Methods       []string
		Credentials   []string
		Headers       []string
		ExposeHeaders []string
		once          sync.Once
	}
	Handler func(conf *Config, c *znet.Context)
)

const (
	SafeHeaders = "Content-Type,Authorization,X-Requested-With,Accept,Origin,Cache-Control,X-File-Name,X-CSRF-Token"
)

func Default() znet.HandlerFunc { _ = "STUB: not implemented"; return *new(znet.HandlerFunc) }

func newAllowOrigins(allowAllHeaders bool) znet.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(znet.HandlerFunc)
}

func AllowAll() znet.HandlerFunc { _ = "STUB: not implemented"; return *new(znet.HandlerFunc) }

func AllowAllOrigins() znet.HandlerFunc { _ = "STUB: not implemented"; return *new(znet.HandlerFunc) }

func NewAllowHeaders() (addAllowHeader func(header string), handler znet.HandlerFunc) {
	_ = "STUB: not implemented"
	return nil, *new(znet.HandlerFunc)
}

func validateConfig(conf *Config) error { _ = "STUB: not implemented"; return nil }

func extractOriginFromReferer(referer string) string { _ = "STUB: not implemented"; return "" }

func (conf *Config) initConfig() { _ = "STUB: not implemented"; return }

// Check if any header is "*" (allow all)

func New(conf *Config) znet.HandlerFunc { _ = "STUB: not implemented"; return *new(znet.HandlerFunc) }

func validateOrigin(origin string) bool { _ = "STUB: not implemented"; return false }

func isOriginAllowed(origin string, conf *Config) bool { _ = "STUB: not implemented"; return false }

func getAllowedHeaders(conf *Config, req *http.Request) string {
	_ = "STUB: not implemented"
	return ""
}

func applyCors(c *znet.Context, conf *Config) bool { _ = "STUB: not implemented"; return false }
