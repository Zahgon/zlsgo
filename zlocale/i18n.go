// Package zlocale provides internationalization (i18n) functionality for Go applications.
// It supports multiple languages, named translation tables, parameterized translations,
// and fallback mechanisms. The module is designed to be lightweight and performant,
// utilizing buffer pools from zutil for efficient string operations.
package zlocale

import (
	"fmt"
	"sync"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zsync"
)

var (
	// ErrLanguageNotFound is returned when a requested language is not available
	ErrLanguageNotFound = fmt.Errorf("language not found")

	// ErrKeyNotFound is returned when a translation key is not found
	ErrKeyNotFound = fmt.Errorf("translation key not found")

	// Global i18n instance for convenience functions
	defaultI18n *I18n
	defaultOnce sync.Once
)

// languageNames maps language codes to their display names
var languageNames = map[string]string{
	"en":    "English",
	"zh-CN": "简体中文",
	"ja":    "日本語",
	"zh":    "简体中文",
}

// Language represents a language with its translation data.
// It provides thread-safe access to translations and cached template processing.
type Language struct {
	// templateCache is the cache system used for template processing
	templateCache TemplateCache
	// data contains the translation key-value pairs
	data      map[string]string
	dataMutex sync.RWMutex
	// mutex protects concurrent access to the translation data using read-biased mutex for performance
	mutex *zsync.RBMutex

	// Legacy template cache for backward compatibility
	templates map[string]*zstring.Template
	// templateMutex protects concurrent access to the legacy templates cache
	templateMutex sync.RWMutex
	// code is the language code (e.g., "en", "zh-CN", "ja")
	code string
	// name is the human-readable language name (e.g., "English", "简体中文")
	name string
	// templateCount tracks the number of cached templates for LRU management
	templateCount int
	// maxTemplates limits the cache size to prevent memory leaks
	maxTemplates int
	// Flag to determine which cache system to use
	useFastCache bool
}

// I18n manages multiple languages and provides translation functionality
type I18n struct {
	// languages contains all loaded languages
	languages map[string]*Language
	// mutex protects concurrent access to the languages map using read-biased mutex for performance
	mutex *zsync.RBMutex
	// defaultLang is the fallback language code
	defaultLang string
	// currentLang is the currently active language code
	currentLang string
}

// New creates a new I18n instance with the specified default language
func New(defaultLang string) *I18n { _ = "STUB: not implemented"; return nil }

// NewDefault creates a new I18n instance with "en" as the default language
func NewDefault() *I18n {
	_ = "STUB: not implemented"

	// getDefault returns the global default i18n instance
	return nil
}

func getDefault() *I18n { _ = "STUB: not implemented"; return nil }

// LoadLanguage loads a language with translation data
func (i *I18n) LoadLanguage(langCode, langName string, data map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadLanguageWithConfig loads a language with translation data and cache configuration
func (i *I18n) LoadLanguageWithConfig(langCode, langName string, data map[string]string, cache TemplateCache) error {
	_ = "STUB: not implemented"
	return nil
}

// SetLanguage sets the current active language
func (i *I18n) SetLanguage(langCode string) error { _ = "STUB: not implemented"; return nil }

// GetLanguage returns the current active language code
func (i *I18n) GetLanguage() string { _ = "STUB: not implemented"; return "" }

// GetLoadedLanguages returns a map of loaded language codes and their names
func (i *I18n) GetLoadedLanguages() map[string]string { _ = "STUB: not implemented"; return nil }

// T translates a key using the current language
func (i *I18n) T(key string, args ...interface{}) string { _ = "STUB: not implemented"; return "" }

// TWithLang translates a key using the specified language
func (i *I18n) TWithLang(langCode, key string, args ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// getTranslationWithLanguage retrieves the translation for a specific language and key, also returns the Language object
func (i *I18n) getTranslationWithLanguage(langCode, key string) (string, *Language) {
	_ = "STUB: not implemented"
	return "", nil
}

// getOrCreateTemplate retrieves or creates a cached template for the given translation string
// Uses the new cache system for improved performance and memory management
func (l *Language) getOrCreateTemplate(templateStr string) (*zstring.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getOrCreateTemplateFastCache uses the new FastCache system for template caching
func (l *Language) getOrCreateTemplateFastCache(templateStr string) (*zstring.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getOrCreateTemplateLegacy uses the original map-based caching for backward compatibility
func (l *Language) getOrCreateTemplateLegacy(templateStr string) (*zstring.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// formatTranslationWithTemplate formats a translation string using optimized template processing
func (i *I18n) formatTranslationWithTemplate(lang *Language, templateStr string, args ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// HasLanguage checks if a language is loaded
func (i *I18n) HasLanguage(langCode string) bool { _ = "STUB: not implemented"; return false }

// HasKey checks if a translation key exists for a specific language
func (i *I18n) HasKey(langCode, key string) bool { _ = "STUB: not implemented"; return false }

// RemoveLanguage removes a language from the i18n instance
func (i *I18n) RemoveLanguage(langCode string) error { _ = "STUB: not implemented"; return nil }

// GetMemoryUsage returns memory usage statistics for monitoring
func (i *I18n) GetMemoryUsage() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// GetCacheStats returns detailed cache statistics for all languages
func (i *I18n) GetCacheStats() map[string]CacheStats { _ = "STUB: not implemented"; return nil }

// ClearTemplateCache clears the template cache for all languages to free memory
func (i *I18n) ClearTemplateCache() { _ = "STUB: not implemented"; return }

// Close closes all cache resources and stops background cleaners
func (i *I18n) Close() { _ = "STUB: not implemented"; return }

// LoadLanguage loads a language into the global i18n instance
func LoadLanguage(langCode, langName string, data map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetLanguage sets the active language for the global i18n instance
func SetLanguage(langCode string) error { _ = "STUB: not implemented"; return nil }

// GetLanguage returns the current active language from the global i18n instance
func GetLanguage() string { _ = "STUB: not implemented"; return "" }

// T translates a key using the global i18n instance
func T(key string, args ...interface{}) string { _ = "STUB: not implemented"; return "" }

// TWithLang translates a key using the specified language in the global i18n instance
func TWithLang(langCode, key string, args ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// HasLanguage checks if a language is loaded in the global i18n instance
func HasLanguage(langCode string) bool { _ = "STUB: not implemented"; return false }

// HasKey checks if a translation key exists in the global i18n instance
func HasKey(langCode, key string) bool { _ = "STUB: not implemented"; return false }

// GetLoadedLanguages returns all loaded languages from the global i18n instance
func GetLoadedLanguages() map[string]string { _ = "STUB: not implemented"; return nil }
