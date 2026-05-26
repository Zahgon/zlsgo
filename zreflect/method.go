package zreflect

import (
	"reflect"
)

// GetAllMethod get all methods of struct
func GetAllMethod(s interface{}, fn func(numMethod int, m reflect.Method) error) error {
	_ = "STUB: not implemented"
	return nil
}

// RunAssignMethod run assign methods of struct
func RunAssignMethod(st interface{}, filter func(methodName string) bool, args ...interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}
