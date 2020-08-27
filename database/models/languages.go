package models

import "database/sql"

// Language - Language for the app.
type Language struct {
	Code   string
	Status string

	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (l Language) GetIdentifier() (identifier interface{}) {
	if l.Code != "" {
		identifier = l.Code
	}

	return
}

// Languages - Alias for a language array.
type Languages []Language
