package ztype

// SliceType is a slice of Type objects that provides helper methods for
// working with collections of values with automatic type conversion.
type SliceType []Type

// Len returns the number of elements in the slice.
func (s SliceType) Len() int {
	_ = "STUB: not implemented"

	// MarshalJSON implements the json.Marshaler interface.
	// It marshals the underlying values rather than the Type wrappers.
	return 0
}

func (s SliceType) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Index returns the element at the specified index.
// Returns an empty Type if the index is out of bounds.
func (s SliceType) Index(i int) Type { _ = "STUB: not implemented"; return *new(Type) }

// Last returns the last element in the slice.
// Returns an empty Type if the slice is empty.
func (s SliceType) Last() Type { _ = "STUB: not implemented"; return *new(Type) }

// First returns the first element in the slice.
// Returns an empty Type if the slice is empty.
func (s SliceType) First() Type {
	_ = "STUB: not implemented"

	// Value returns the underlying values as a slice of interface{}.
	// This unwraps all Type objects to their original values.
	return *new(Type)
}

func (s SliceType) Value() []interface{} { _ = "STUB: not implemented"; return nil }

// String converts all elements in the slice to strings and returns them as a []string.
// Each element is converted using the Type.String() method.
func (s SliceType) String() []string { _ = "STUB: not implemented"; return nil }

// Int converts all elements in the slice to integers and returns them as a []int.
// Each element is converted using the Type.Int() method.
func (s SliceType) Int() []int { _ = "STUB: not implemented"; return nil }

// Maps converts all elements in the slice to Map objects and returns them as a Maps slice.
// Each element is converted using the Type.Map() method.
func (s SliceType) Maps() Maps { _ = "STUB: not implemented"; return *new(Maps) }

// func (s SliceType) Slice() []SliceType {
// 	ss := make([]SliceType, 0, len(s))
// 	for i := range s {
// 		ss = append(ss, s[i].Slice())
// 	}
// 	return ss
// }

// Slice converts a value to a SliceType.
// Deprecated: please use ToSlice instead.
func Slice(value interface{}, noConv ...bool) SliceType {
	_ = "STUB: not implemented"
	return *new(SliceType)
}

// SliceStrToAny converts a slice of strings to a slice of interface{} values.
// This is useful when you need to pass a string slice to a function that expects interface{} values.
func SliceStrToAny(slice []string) []interface{} { _ = "STUB: not implemented"; return nil }

// ToSlice converts various types to a SliceType.
// If noConv is true, it will not attempt to convert non-slice values (like strings) to slices.
// Handles []interface{}, []string, []int, []int64, and can parse JSON strings into slices.
func ToSlice(value interface{}, noConv ...bool) (s SliceType) {
	_ = "STUB: not implemented"
	return *new(SliceType)
}
