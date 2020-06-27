package models

import "database/sql"

// User - User for the app.
type User struct {
	ID         int `json:"id,omitempty"`
	LanguageID int `json:"languageId,omitempty"`

	Username string       `json:"username,omitempty"`
	Password string       `json:"password,omitempty"`
	Name     string       `json:"name,omitempty"`
	LastName string       `json:"lastName,omitempty"`
	Birthday sql.NullTime `json:"birthday,omitempty"`
	Theme    string       `json:"theme,omitempty"`
}

// Users - Alias for a user array.
type Users []User
