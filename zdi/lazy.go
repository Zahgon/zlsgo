package zdi

import (
	"reflect"
)

// Provide registers a provider function with the injector.
// A provider is a function that, when invoked, returns one or more values to be injected.
func (inj *injector) Provide(provider interface{}, opt ...Option) (override []reflect.Type) {
	_ = "STUB: not implemented"
	return nil
}
