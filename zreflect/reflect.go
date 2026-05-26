package zreflect

import (
	"reflect"
	"unsafe"
)

type (
	Type  = *rtype
	rtype struct{}
	flag  uintptr
	Value struct {
		typ Type
		ptr unsafe.Pointer
		flag
	}
)

// GetUnexportedField retrieves the value of an unexported field from a struct.
// This is a hazardous operation that bypasses Go's type safety and should be used with extreme caution.
//
// v is the reflect.Value of the struct containing the unexported field.
// field is the name of the unexported field to access.
//
// It returns the value of the unexported field, or an error if the field
// doesn't exist or cannot be accessed.
func GetUnexportedField(v reflect.Value, field string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetUnexportedField sets the value of an unexported field in a struct.
// This is a hazardous operation that bypasses Go's type safety and should be used with extreme caution.
//
// v is the reflect.Value of the struct containing the unexported field.
// field is the name of the unexported field to modify.
// value is the new value to set for the field.
//
// It returns an error if the field doesn't exist, cannot be modified, or if the value type doesn't match.
func SetUnexportedField(v reflect.Value, field string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func canAssignNil(kind reflect.Kind) bool { _ = "STUB: not implemented"; return false }

// getField is an internal helper function that retrieves a field from a struct by name.
// It returns the field's reflect.Value, a boolean indicating if the field is exported,
// and an error if the field doesn't exist or cannot be accessed.
func getField(v reflect.Value, field string) (reflect.Value, bool, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false, nil
}
