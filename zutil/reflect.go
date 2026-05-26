package zutil

import (
	"reflect"
)

func SetValue(vTypeOf reflect.Kind, vValueOf reflect.Value, value interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// setStruct todo unfinished
func setStruct(v reflect.Value, value interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func ReflectStructField(v reflect.Type, fn func(
	numField int, fieldTag string, field reflect.StructField) error, tag ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func ReflectForNumField(v reflect.Value, fn func(fieldName, fieldTag string,
	kind reflect.Kind, field reflect.Value) error, tag ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}

//  && tfield.Anonymous

// GetAllMethod get all methods of struct
func GetAllMethod(s interface{}, fn func(numMethod int, m reflect.Method) error) error {
	_ = "STUB: not implemented"
	return nil
}

// RunAllMethod run all methods of struct
func RunAllMethod(st interface{}, args ...interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RunAssignMethod run assign methods of struct
func RunAssignMethod(st interface{}, filter func(methodName string) bool, args ...interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}
