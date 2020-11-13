package models

import (
	"regexp"
	"time"

	"github.com/coffemanfp/beppin/utils"
)

// User - User for the app.
type User struct {
	ID       int64  `json:"id,omitempty"`
	Language string `json:"language,omitempty"`

	AvatarURL string     `json:"avatarURL,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email,omitempty"`
	Name      string     `json:"name,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Theme     string     `json:"theme,omitempty"`
	Currency  string     `json:"currency,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Users - Alias for a user array.
type Users []User

// Validate - Validates a user.
func (u User) Validate(action string) (valid bool) {
	switch action {
	case "login":
		valid = u.validateLogin()
	case "signup":
		valid = u.validateSignup()
	}

	return
}

// validateLogin - Validates a user login.
func (u User) validateLogin() (valid bool) {
	valid = true

	if u.Password == "" {
		valid = false
		return
	}

	// If it is not equal to any
	if !(utils.ValidateEmail(u.Email) || u.ValidateUsername()) {
		valid = false
	}
	return
}

// validateSignup - Validates a user signup.
func (u User) validateSignup() (valid bool) {
	valid = true

	if u.Username == "" ||
		u.Email == "" ||
		u.Password == "" {
		valid = false
	}

	return
}

// ValidateUsername - Validate a username.
func (u User) ValidateUsername() (valid bool) {
	if u.Username == "" {
		return
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`)

	valid = re.MatchString(u.Username)
	return
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
