// Package zvalid data verification.
package zvalid

import (
	"errors"

	"github.com/sohaha/zlsgo/zjson"
)

type (
	// Engine valid engine
	Engine struct {
		err          error
		defaultValue interface{}
		name         string
		value        string
		sep          string
		queue        []queueT
		valueInt     int
		valueFloat   float64
		setRawValue  bool
		silent       bool
		result       bool
	}
	queueT func(v *Engine) *Engine
)

// ErrNoValidationValueSet no verification value set
var ErrNoValidationValueSet = errors.New("未设置验证值")

// New valid
func New() Engine { _ = "STUB: not implemented"; return *new(Engine) }

// Int use int new valid
func Int(value int, name ...string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

// Text use int new valid
func Text(value string, name ...string) Engine { _ = "STUB: not implemented"; return *new(Engine) }

func JSON(json *zjson.Res, rules map[string]Engine) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Required Must have a value (zero values ​​other than "" are allowed). If this rule is not used, when the parameter value is "", data validation does not take effect by default
func (v Engine) Required(customError ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// Customize customize valid
func (v Engine) Customize(fn func(rawValue string, err error) (newValue string, newErr error)) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

func pushQueue(v *Engine, fn queueT, DisableCheckErr ...bool) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

func ignore(v *Engine) bool { _ = "STUB: not implemented"; return false }

func notEmpty(v *Engine) bool { _ = "STUB: not implemented"; return false }

func setError(v *Engine, msg string, customError ...string) error {
	_ = "STUB: not implemented"
	return nil
}
