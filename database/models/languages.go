package models

import "database/sql"

// Language - Language for the app.
type Language struct {
	ID     int64
	Code   string
	Status string

	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
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

// Languages - Alias for a language array.
type Languages []Language
