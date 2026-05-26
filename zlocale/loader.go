package zlocale

// Loader provides functionality to load translations from maps
// Simplified version that only supports embedded translations
type Loader struct {
	i18n *I18n
}

// NewLoader creates a new loader for the given i18n instance
func NewLoader(i18n *I18n) *Loader { _ = "STUB: not implemented"; return nil }

// LoadTranslationsFromMap loads translations from a map[string]string
func (l *Loader) LoadTranslationsFromMap(langCode, langName string, translations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddCustomTranslation adds a custom translation to an existing language
// This allows runtime extension of the built-in translations
func (l *Loader) AddCustomTranslation(langCode, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAvailableLanguages returns the list of available language codes
func (l *Loader) GetAvailableLanguages() []string { _ = "STUB: not implemented"; return nil }

// LoadTranslationsFromMap loads translations from a map using the global i18n instance
func LoadTranslationsFromMap(langCode, langName string, translations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddCustomTranslation adds a custom translation using the global i18n instance
func AddCustomTranslation(langCode, key, value string) error { _ = "STUB: not implemented"; return nil }
