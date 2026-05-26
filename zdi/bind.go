package zdi

// Resolve resolves dependencies for the given pointers.
// It iterates through each pointer, determines its underlying type,
// and injects the corresponding value from the injector.
// Returns an error if a dependency cannot be found or if a value cannot be set.
func (inj *injector) Resolve(v ...Pointer) error { _ = "STUB: not implemented"; return nil }

// Apply injects dependencies into the fields of a struct or sets a pointer value.
// If the provided pointer is a struct, it iterates through its fields.
// For fields tagged with `di`, it resolves and injects the corresponding dependency.
// If the pointer is not a struct, it attempts to resolve and set the value directly.
// Returns an error if a dependency cannot be found for a tagged field or if a value cannot be set.
func (inj *injector) Apply(p Pointer) error { _ = "STUB: not implemented"; return nil }
