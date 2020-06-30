package models

import (
	"regexp"
	"time"
)

// User - User for the app.
type User struct {
	ID       int    `json:"id,omitempty"`
	Language string `json:"language,omitempty"`

	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	Name      string     `json:"name,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Theme     string     `json:"theme,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
}

// Users - Alias for a user array.
type Users []User

// Validate - Validates a user.
func (u User) Validate() (valid bool) {
	valid = true

	switch "" {
	case u.Username:
	case u.Password:
	case u.Name:
	case u.LastName:
		valid = false
	}

	if u.Birthday == nil || u.Birthday.IsZero() {
		valid = false
	}

	return
}

// ValidateUsername - Validate a username.
func (u User) ValidateUsername() (valid bool) {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)

	valid = re.MatchString(u.Username)
	return
}
