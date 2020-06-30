package models

import (
	"time"

	"github.com/coffemanfp/beppin-server/utils"
)

// Language status.
const (
	LanguageUnavailable = "unavailable" // LanguageUnavailable - Existing language but unavailable.
	LanguageInProgress  = "in-progress" // LanguageInProgress - Developing language.
	LanguageAvailable   = "available"   // LanguageAvailable - AvailableLanguage
)

// Language - Language for the app.
type Language struct {
	ID int `json:"id,omitempty"`

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

	switch l.Status {
	case LanguageUnavailable:
	case LanguageInProgress:
	case LanguageAvailable:
	default:
		valid = false
	}

	valid = utils.ValidateLanguageCode(l.Code)
	return
}
