package models

// Language status.
const (
	LanguageUnavailable = iota + 1 // LanguageUnavailable - Existing language but unavailable.
	LanguageInProgress             // LanguageInProgress - Developing language.
	LanguageAvailable              // LanguageAvailable - AvailableLanguage
)

// Language - Language for the app.
type Language struct {
	ID int `json:"id,omitempty"`

	Code   string
	Status int
}

// Languages - Alias for a language array.
type Languages []Language
