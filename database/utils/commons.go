package utils

import (
	"database/sql"

	"github.com/coffemanfp/beppin/models"
)

// Helper types for null database retorning

type (
	nullUserData struct {
		AvatarURL *sql.NullString
		Name      *sql.NullString
		LastName  *sql.NullString
		Birthday  *sql.NullTime
		UpdatedAt *sql.NullTime
	}
	nullProductData struct {
		UpdatedAt *sql.NullTime
	}
	nullLanguageData struct {
		UpdatedAt *sql.NullTime
	}
)

// Fills the fields data if isn't empty

func (n nullUserData) setResults(user *models.User) {
	if n.AvatarURL != nil {
		user.AvatarURL = n.AvatarURL.String
	}
	if n.Name != nil {
		user.Name = n.Name.String
	}
	if n.LastName != nil {
		user.LastName = n.LastName.String
	}
	if n.Birthday != nil {
		user.Birthday = &n.Birthday.Time
	}
	if n.UpdatedAt != nil {
		user.UpdatedAt = &n.UpdatedAt.Time
	}
}

func (n nullProductData) setResults(product *models.Product) {
	if n.UpdatedAt != nil {
		product.UpdatedAt = &n.UpdatedAt.Time
	}
}

func (n nullLanguageData) setResults(language *models.Language) {
	if n.UpdatedAt != nil {
		language.UpdatedAt = &n.UpdatedAt.Time
	}
}
