package models

import "database/sql"

// Language - Language for the app.
type Language struct {
	ID int

	Code   string
	Status string

	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
}

// Languages - Alias for a language array.
type Languages []Language
