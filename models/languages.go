package models

import (
	"regexp"
	"time"
)

// Language status.
const (
	LanguageUnavailable = "unavailable" // LanguageUnavailable - Existing language but unavailable.
	LanguageInProgress  = "in-progress" // LanguageInProgress - Developing language.
	LanguageAvailable   = "available"   // LanguageAvailable - AvailableLanguage
)

// Language - Language for the app.
type Language struct {
	ID int64 `json:"id,omitempty"`

	Code   string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Languages - Alias for a language array.
type Languages []Language

// Validate - Validate a language.
func (l Language) Validate() (valid bool) {
	valid = true

	// If it is not equal to any
	if !(l.Status == LanguageUnavailable ||
		l.Status == LanguageInProgress ||
		l.Status == LanguageAvailable) {
		valid = false
	}

	valid = l.ValidateCode()
	return
}

// ValidateCode - Validates the language code.
func (l Language) ValidateCode() (valid bool) {
	rx := regexp.MustCompile(`^[a-z]{2,2}-[A-Z]{2,2}$`)

	valid = rx.MatchString(l.Code)
	return
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (l Language) GetIdentifier() (identifier interface{}) {
	if l.ID != 0 {
		identifier = l.ID
	} else if l.Code != "" {
		identifier = l.Code
	}

	return
}
