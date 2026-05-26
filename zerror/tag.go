package zerror

type TagKind string

const (
	// Empty error
	None TagKind = ""
	// Internal errors, This means that some invariants expected by the underlying system have been broken
	Internal TagKind = "INTERNAL"
	// The operation was cancelled, typically by the caller
	Cancelled TagKind = "CANCELLED"
	// The client specified an invalid argument
	InvalidInput TagKind = "INVALID_INPUT"
	// Some requested entity was not found
	NotFound TagKind = "NOT_FOUND"
	// The caller does not have permission to execute the specified operation
	PermissionDenied TagKind = "PERMISSION_DENIED"
	// The request does not have valid authentication credentials for the operation
	Unauthorized TagKind = "UNAUTHORIZED"
)

func (t TagKind) Wrap(err error, text string) error { _ = "STUB: not implemented"; return nil }

func (t TagKind) Text(text string) error { _ = "STUB: not implemented"; return nil }

type withTag struct {
	wrapErr error
	tag     TagKind
}

func (e *withTag) Error() string { _ = "STUB: not implemented"; return "" }

func WrapTag(tag TagKind) External { _ = "STUB: not implemented"; return *new(External) }

func GetTag(err error) TagKind { _ = "STUB: not implemented"; return *new(TagKind) }
