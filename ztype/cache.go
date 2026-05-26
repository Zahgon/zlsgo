package ztype

import (
	"reflect"
	"sync"
)

// fieldInfo cache struct field info
type fieldInfo struct {
	Type       reflect.Type
	Name       string
	Tag        string
	Options    []string
	Index      int
	IsTime     bool
	IsExported bool
}

// structCacheEntry struct cache entry
type structCacheEntry struct {
	TypeName string
	Fields   []fieldInfo
}

// structCache struct cache using reflect.Type as key to avoid name conflicts
var structCache sync.Map

// getStructInfo get struct info, priority from cache
func getStructInfo(t reflect.Type) []fieldInfo { _ = "STUB: not implemented"; return nil }

// parseStructFields parse struct fields
func parseStructFields(t reflect.Type) []fieldInfo { _ = "STUB: not implemented"; return nil }

// parseTagOptions parse tag options
func parseTagOptions(opt string) []string { _ = "STUB: not implemented"; return nil }

// hasOption check field has specific option
func (f *fieldInfo) hasOption(option string) bool { _ = "STUB: not implemented"; return false }
