package zvalid

// Ok not err
func (v Engine) Ok() bool { _ = "STUB: not implemented"; return false }

// Error or whether the verification fails
func (v Engine) Error() error { _ = "STUB: not implemented"; return nil }

// Value get the final value
func (v Engine) Value() (value string) { _ = "STUB: not implemented"; return "" }

// String to string
func (v Engine) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Bool to bool
func (v Engine) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Int convert to int
func (v Engine) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Float64 convert to float64
func (v Engine) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Split converted to [] string
func (v Engine) Split(sep string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Valid executes the validation queue and returns the final result.
// Note: All queue operations will be executed even after validation errors occur,
// ensuring that operations like Default() can work properly.
// This method includes defensive checks to prevent panic from nil queue functions.
func (v *Engine) valid() *Engine { _ = "STUB: not implemented"; return nil }

// Skip nil queue functions to prevent panic

// Skip nil results to prevent nil pointer dereference

// Continue executing queue items even after error to allow Default() to work

// SetAlias set alias
func (v Engine) SetAlias(name string) Engine {
	_ = "STUB: not implemented"
	return *

	// Verifi validate specified data
	new(Engine)
}

func (v Engine) Verifi(value string, name ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// VerifiAny validate specified data
func (v Engine) VerifiAny(value interface{}, name ...string) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

// v.err = setError(&v, "unsupported type")
