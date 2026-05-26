package zreflect

import (
	"errors"
	"reflect"
)

// ForEachMethod iterates through all methods of a given value and calls the provided function for each method.
// This simplifies the process of examining and working with methods via reflection.
//
// valof is the reflect.Value whose methods will be iterated.
// fn is a callback function that receives the method index, method information, and method value.
//
// It returns any error returned by the callback function, or nil if all iterations complete successfully.
func ForEachMethod(valof reflect.Value, fn func(index int, method reflect.Method, value reflect.Value) error) error {
	_ = "STUB: not implemented"
	return nil
}

// SkipChild is a special error value that can be returned from ForEach and ForEachValue callbacks
// to indicate that the current struct field's children should be skipped during iteration.
var SkipChild = errors.New("skip struct")

// ForEach iterates through all fields of a struct type recursively and calls the provided function for each field.
// This provides a convenient way to examine and process struct fields via reflection.
//
// typ is the reflect.Type of the struct to iterate through.
// fn is a callback function that receives the parent path, field index, tag value, and field information.
//
// It returns any error returned by the callback function, or nil if all iterations complete successfully.
func ForEach(typ reflect.Type, fn func(parent []string, index int, tag string, field reflect.StructField) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ForEachValue iterates through all fields of a struct value recursively and calls the provided function for each field.
// Similar to ForEach, but also provides access to the field's value.
//
// val is the reflect.Value of the struct to iterate through.
// fn is a callback function that receives the parent path, field index, tag value, field information, and field value.
//
// It returns any error returned by the callback function, or nil if all iterations complete successfully.
func ForEachValue(val reflect.Value, fn func(parent []string, index int, tag string, field reflect.StructField, val reflect.Value) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}
