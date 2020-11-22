package utils

import (
	"database/sql"

	"github.com/coffemanfp/beppin/models"
)

// Helper types for null database retorning

type (
	nullUserData struct {
		AvatarID   *sql.NullInt64
		AvatarPath *sql.NullString
		Name       *sql.NullString
		LastName   *sql.NullString
		Birthday   *sql.NullTime
		UpdatedAt  *sql.NullTime
	}
	nullProductData struct {
		Description *sql.NullString
		UpdatedAt   *sql.NullTime
	}
	nullLanguageData struct {
		UpdatedAt *sql.NullTime
	}
	nullFileData struct {
		ID        *sql.NullInt64
		Path      *sql.NullString
		CreatedAt *sql.NullTime
		UpdatedAt *sql.NullTime
	}
)

// Fills the fields data if isn't empty

func (n nullUserData) setResults(user *models.User) {
	if n.AvatarID != nil {
		if user.Avatar == nil {
			user.Avatar = new(models.File)
		}
		user.Avatar.ID = n.AvatarID.Int64
	}
	if n.AvatarPath != nil {
		if user.Avatar == nil {
			user.Avatar = new(models.File)
		}
		user.Avatar.Path = n.AvatarPath.String
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
	if n.Description != nil {
		product.Description = n.Description.String
	}
	if n.UpdatedAt != nil {
		product.UpdatedAt = &n.UpdatedAt.Time
	}
}

func (n nullLanguageData) setResults(language *models.Language) {
	if n.UpdatedAt != nil {
		language.UpdatedAt = &n.UpdatedAt.Time
	}
}

func (n nullFileData) setResults(file *models.File) {
	if n.UpdatedAt != nil {
		file.UpdatedAt = &n.UpdatedAt.Time
	}
}
