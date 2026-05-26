package ztype

import (
	"reflect"
)

type (
	// StruBuilder provides functionality to dynamically build struct types at runtime.
	// It supports creating regular structs, map[T]struct, and []struct types.
	StruBuilder struct {
		key       reflect.Type
		fields    map[string]*StruField
		fieldKeys []string
		typ       int
	}
	// StruField represents a field in a dynamically built struct.
	// It contains the field type and tag information.
	StruField struct {
		typ interface{}
		tag string
	}
)

const (
	typeStruct = iota
	typeMapStruct
	typeSliceStruct
)

// NewStructFromValue creates a new StruBuilder from an existing struct value.
// It analyzes the provided struct and copies all its fields to the builder.
func NewStructFromValue(v interface{}) (*StruBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewStruct creates a new StruBuilder for building regular struct types.
func NewStruct() *StruBuilder { _ = "STUB: not implemented"; return nil }

// NewMapStruct creates a new StruBuilder for building map[T]struct types.
// The key parameter specifies the map key type.
func NewMapStruct(key interface{}) *StruBuilder { _ = "STUB: not implemented"; return nil }

// NewSliceStruct creates a new StruBuilder for building []struct types.
func NewSliceStruct() *StruBuilder { _ = "STUB: not implemented"; return nil }

// Copy copies the configuration from another StruBuilder while preserving the current type.
func (b *StruBuilder) Copy(v *StruBuilder) *StruBuilder { _ = "STUB: not implemented"; return nil }

func (b *StruBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Merge merges fields from one or more struct values into this builder.
// All provided values must be struct types.
func (b *StruBuilder) Merge(values ...interface{}) error { _ = "STUB: not implemented"; return nil }

// func (b *StruBuilder) AddFunc(name string, fieldType interface{}, tag ...string) *StruBuilder {
// 	reflect.MakeFunc()
// 	return b
// }

// AddField adds a new field to the struct being built.
// The fieldType can be a reflect.Type, another StruBuilder, or any value whose type will be used.
// Optional tag strings will be joined with spaces to form the struct tag.
func (b *StruBuilder) AddField(name string, fieldType interface{}, tag ...string) *StruBuilder {
	_ = "STUB: not implemented"
	return nil
}

// RemoveField removes a field from the struct being built.
func (b *StruBuilder) RemoveField(name string) *StruBuilder { _ = "STUB: not implemented"; return nil }

// HasField checks if a field with the given name exists in the struct.
func (b *StruBuilder) HasField(name string) bool { _ = "STUB: not implemented"; return false }

// GetField retrieves a field by name. Returns nil if the field doesn't exist.
func (b *StruBuilder) GetField(name string) *StruField { _ = "STUB: not implemented"; return nil }

// FieldNames returns a slice containing all field names in the struct.
func (b *StruBuilder) FieldNames() []string {
	_ = "STUB: not implemented"

	// Interface returns the built struct as an interface{}.
	return nil
}

func (b *StruBuilder) Interface() interface{} { _ = "STUB: not implemented"; return nil }

// Type returns the reflect.Type of the built struct.
func (b *StruBuilder) Type() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// Value returns a new reflect.Value of the built struct type.
func (b *StruBuilder) Value() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

// SetType sets the type of this field.
func (f *StruField) SetType(typ interface{}) *StruField { _ = "STUB: not implemented"; return nil }

// SetTag sets the struct tag for this field.
func (f *StruField) SetTag(tag string) *StruField { _ = "STUB: not implemented"; return nil }
