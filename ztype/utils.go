package ztype

import (
	"reflect"
	"time"
)

// GetType Get variable type
func GetType(s interface{}) string { _ = "STUB: not implemented"; return "" }

func reflectPtr(r reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func parsePath(path string, v interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

var timeType = reflect.TypeOf(time.Time{})
