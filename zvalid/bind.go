package zvalid

import (
	"reflect"
)

// ValidEle ValidEle
type ValidEle struct {
	target interface{}
	source Engine
}

// Silent an error occurred during filtering, no error is returned
func (v Engine) Silent() Engine {
	_ = "STUB: not implemented"
	return *

	// Default if a filtering error occurs, the default value is assigned to the variable
	new(Engine)
}

func (v Engine) Default(value interface{}) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// Separator specify the separator of the slice type
func (v Engine) Separator(sep string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// BatchError  multiple error
func BatchError(rules ...Engine) error { _ = "STUB: not implemented"; return nil }

// Batch assign multiple filtered results to the specified object
func Batch(elements ...*ValidEle) error { _ = "STUB: not implemented"; return nil }

// BatchVar assign the filtered result to the specified variable
func BatchVar(target interface{}, source Engine) *ValidEle { _ = "STUB: not implemented"; return nil }

// Var assign the filtered result to the specified variable
func Var(target interface{}, source Engine, name ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func setRawValue(k reflect.Kind, val reflect.Value, value string, sep string) error {
	_ = "STUB: not implemented"
	return nil
}

func setDefaultValue(targetTypeOf reflect.Kind, targetValueOf reflect.Value, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
