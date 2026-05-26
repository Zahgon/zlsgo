package zstring

import (
	"io"
)

// Template implements a simple template engine that replaces tags in a template string.
// It supports custom start and end tag delimiters and efficient processing.
type Template struct {
	template string // The original template string
	startTag []byte // Byte representation of the opening tag delimiter
	endTag   []byte // Byte representation of the closing tag delimiter

	texts [][]byte // Slices of text between tags
	tags  []string // The tag names extracted from the template
}

// NewTemplate creates a new template with the specified template string and tag delimiters.
// It parses the template and returns an error if the template format is invalid.
func NewTemplate(template, startTag, endTag string) (*Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResetTemplate changes the template string and re-parses it.
// Returns an error if the template format is invalid (e.g., missing end tags).
func (t *Template) ResetTemplate(template string) error { _ = "STUB: not implemented"; return nil }

// Process executes the template, writing the result to the provided writer.
// For each tag encountered, it calls the provided function with the tag name.
// Returns the number of bytes written and any error encountered.
func (t *Template) Process(w io.Writer, fn func(w io.Writer, tag string) (int, error)) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
