package zreflect

import (
	"reflect"
	_ "unsafe"
)

// TypeOf returns the reflection Type of the value v.
// This is similar to reflect.TypeOf but with additional handling for zreflect.Type and zreflect.Value types.
func TypeOf(v interface{}) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

//go:linkname toRType reflect.toType
//go:noescape
func toRType(Type) reflect.Type

// NewType creates a new Type from the given value.
// It handles various input types including Type, Value, reflect.Type, reflect.Value,
// or any other value, converting them to the internal Type representation.
func NewType(v interface{}) Type { _ = "STUB: not implemented"; return *new(Type) }

// Native converts the internal Type representation to the standard reflect.Type.
// This allows interoperability with the standard reflect package.
func (t *rtype) Native() reflect.Type {
	_ = "STUB: not implemented"

	// rtypeToType converts a reflect.Type to the internal Type representation.
	// This is an internal helper function used for type conversions.
	return *new(reflect.Type)
}

func rtypeToType(t reflect.Type) Type { _ = "STUB: not implemented"; return *new(Type) }

//go:linkname typeNumMethod reflect.(*rtype).NumMethod
//go:noescape
func typeNumMethod(Type) int

// NumMethod returns the number of exported methods in the type's method set.
// This is equivalent to reflect.Type.NumMethod() but works on the internal Type representation.
func (t *rtype) NumMethod() int { _ = "STUB: not implemented"; return 0 }
