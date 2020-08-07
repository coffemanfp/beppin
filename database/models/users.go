package models

import "database/sql"

// User - User for the app.
type User struct {
	ID       int
	Language Language

	Username  string
	Password  string
	Name      string
	Email     string
	LastName  string
	Birthday  *sql.NullTime
	Theme     string
	CreatedAt *sql.NullTime
	UpdatedAt *sql.NullTime
}

// GetIdentifier gets the first unique identifier it finds in order of importance.
func (u User) GetIdentifier() (identifier interface{}) {

	if u.ID != 0 {
		identifier = u.ID
	} else if u.Username != "" {
		identifier = u.Username
	} else if u.Email != "" {
		identifier = u.Email
	}

	return
}

// Users - Alias for a user array.
type Users []User
