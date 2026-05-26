package zpool

import (
	"reflect"

	"github.com/sohaha/zlsgo/zdi"
)

type (
	invokerPre func() error
)

func (i invokerPre) Invoke(_ []interface{}) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	_ zdi.PreInvoker = (*invokerPre)(nil)
)

func invokeHandler(v []reflect.Value, err error) error { _ = "STUB: not implemented"; return nil }

func (wp *WorkPool) Injector() zdi.TypeMapper {
	_ = "STUB: not implemented"
	return *new(zdi.TypeMapper)
}

func (wp *WorkPool) handlerFunc(h Task) (fn taskfn) { _ = "STUB: not implemented"; return *new(taskfn) }
