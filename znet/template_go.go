package znet

import (
	"html/template"
	"io"
	"sync"

	"github.com/sohaha/zlsgo/zlog"
)

type htmlEngine struct {
	log       *zlog.Logger
	funcmap   map[string]interface{}
	Templates *template.Template
	directory string
	options   TemplateOptions
	mutex     sync.RWMutex
	loaded    bool
}

type TemplateOptions struct {
	Extension  string
	Layout     string
	DelimLeft  string
	DelimRight string
	Reload     bool
	Debug      bool
}

func getTemplateOptions(debug bool, opt ...func(o *TemplateOptions)) TemplateOptions {
	_ = "STUB: not implemented"
	return *new(TemplateOptions)
}

var _ Template = &htmlEngine{}

func newGoTemplate(e *Engine, directory string, opt ...func(o *TemplateOptions)) *htmlEngine {
	_ = "STUB: not implemented"
	return nil
}

func (e *htmlEngine) AddFunc(name string, fn interface{}) *htmlEngine {
	_ = "STUB: not implemented"
	return nil
}

func (e *htmlEngine) SetFuncMap(m map[string]interface{}) *htmlEngine {
	_ = "STUB: not implemented"
	return nil
}

func (e *htmlEngine) Load() error { _ = "STUB: not implemented"; return nil }

func (e *htmlEngine) Render(out io.Writer, template string, data interface{}, layout ...string) error {
	_ = "STUB: not implemented"
	return nil
}
