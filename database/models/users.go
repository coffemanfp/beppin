package models

import "database/sql"

// User - User for the app.
type User struct {
	ID       int
	Language Language

	Username  string
	Password  string
	Name      string
	LastName  string
	Birthday  *sql.NullTime
	Theme     string
	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
}

// Users - Alias for a user array.
type Users []User
