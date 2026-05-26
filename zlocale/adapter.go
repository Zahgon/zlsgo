package zlocale

import (
	"github.com/sohaha/zlsgo/zsync"
)

// LegacyCacheAdapter implements TemplateCache interface using the original map-based approach
// This maintains backward compatibility and provides a fallback option
type LegacyCacheAdapter struct {
	cache     map[string]*TemplateCacheEntry
	hitCount  int64
	missCount int64
	maxSize   int
	mutex     *zsync.RBMutex
}

// NewLegacyCacheAdapter creates a new legacy map-based cache adapter
func NewLegacyCacheAdapter(maxSize int) *LegacyCacheAdapter { _ = "STUB: not implemented"; return nil }

// Get retrieves a cached template entry by key
func (l *LegacyCacheAdapter) Get(key string) (*TemplateCacheEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set stores a template entry with optional expiration
func (l *LegacyCacheAdapter) Set(key string, entry *TemplateCacheEntry) {
	_ = "STUB: not implemented"
	return
}

// Delete removes a template from the cache
func (l *LegacyCacheAdapter) Delete(key string) { _ = "STUB: not implemented"; return }

// Clear removes all templates from the cache
func (l *LegacyCacheAdapter) Clear() { _ = "STUB: not implemented"; return }

// Count returns the number of cached templates
func (l *LegacyCacheAdapter) Count() int { _ = "STUB: not implemented"; return 0 }

// Stats returns cache statistics
func (l *LegacyCacheAdapter) Stats() CacheStats { _ = "STUB: not implemented"; return *new(CacheStats) }

// Legacy adapter doesn't track evictions

// Rough estimate

// Close cleans up cache resources
func (l *LegacyCacheAdapter) Close() { _ = "STUB: not implemented"; return }
