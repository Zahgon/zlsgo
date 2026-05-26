package zdi

import (
	"reflect"
)

// InvokeWithErrorOnly calls the given function and returns only an error, if any.
// It's a convenience wrapper around Invoke for functions that primarily signal success/failure via an error.
func (inj *injector) InvokeWithErrorOnly(f interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Invoke calls the given function, injecting its dependencies.
// It resolves the function's arguments using the injector's known types.
// Returns the function's return values and an error if dependency resolution or function call fails.
func (inj *injector) Invoke(f interface{}) (values []reflect.Value, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// call is an internal helper to invoke a regular function.
// It resolves dependencies for the function's arguments and calls the function.
func (inj *injector) call(f interface{}, t reflect.Type, numIn int) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Map maps a value to its own type or to a specified interface type within the injector.
// Options can be provided, e.g., WithInterface, to specify the mapping key.
// Returns the type key if it overrides an existing mapping.
func (inj *injector) Map(val interface{}, opt ...Option) (override reflect.Type) {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// Maps maps multiple values into the injector.
// It's a convenience function for calling Map for each provided value.
// Returns a slice of types for which existing mappings were overridden.
func (inj *injector) Maps(values ...interface{}) (override []reflect.Type) {
	_ = "STUB: not implemented"
	return nil
}

// Set maps a reflect.Value to a reflect.Type in the injector.
// This is a lower-level way to directly insert values into the injector's store.
func (inj *injector) Set(typ reflect.Type, val reflect.Value) { _ = "STUB: not implemented"; return }

// Get retrieves a value of the given type from the injector.
// It first checks for directly mapped values. If not found, it checks providers.
// If still not found and the type is an interface, it searches for compatible concrete types.
// Finally, it consults the parent injector, if one exists.
// Returns the reflect.Value and a boolean indicating if the value was found.
func (inj *injector) Get(t reflect.Type) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func providerKey(provider reflect.Value, t reflect.Type) string {
	_ = "STUB: not implemented"
	return ""
}
