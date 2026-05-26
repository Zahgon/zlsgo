package znet

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
)

var preInvokers = make([]reflect.Type, 0)

func registerPreInvoker(preInvokers []reflect.Type, invoker ...zdi.PreInvoker) ([]reflect.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RegisterRender Register Render
func RegisterRender(invoker ...zdi.PreInvoker) (err error) { _ = "STUB: not implemented"; return nil }

// RegisterRender Register Render
func (e *Engine) RegisterRender(invoker ...zdi.PreInvoker) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// BindStruct Bind Struct
func (e *Engine) BindStruct(prefix string, s interface{}, handle ...Handler) error {
	_ = "STUB: not implemented"
	return nil
}
