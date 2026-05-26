package zdi

import (
	"reflect"
)

type PreInvoker interface {
	Invoke([]interface{}) ([]reflect.Value, error)
}

// IsPreInvoker checks if the given handler implements the PreInvoker interface.
// PreInvoker allows for a potentially faster invocation path by bypassing some reflection.
func IsPreInvoker(handler interface{}) bool { _ = "STUB: not implemented"; return false }

// fast is an internal helper for invoking a PreInvoker.
// It resolves dependencies for the PreInvoker's arguments as interface{} slices
// and then calls its Invoke method.
func (inj *injector) fast(f PreInvoker, t reflect.Type, numIn int) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
