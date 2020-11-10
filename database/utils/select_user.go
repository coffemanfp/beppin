package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectUser - Selects a user.
func SelectUser(db *sql.DB, userToFind models.User) (user models.User, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := userToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, language, avatar, username, email, name, last_name, birthday, theme, currency, created_at, updated_at
		FROM
			users
		WHERE
			id = $1 OR username = $2 OR email = $3
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	// Helper value for null database retorning
	nullData := struct {
		AvatarURL *sql.NullString
		Name      *sql.NullString
		LastName  *sql.NullString
		Birthday  *sql.NullTime
		UpdatedAt *sql.NullTime
	}{}

	err = stmt.QueryRow(
		userToFind.ID,
		userToFind.Username,
		userToFind.Email,
	).Scan(
		&user.ID,
		&user.Language,
		&nullData.AvatarURL,
		&user.Username,
		&user.Email,
		&nullData.Name,
		&nullData.LastName,
		&nullData.Birthday,
		&user.Theme,
		&user.Currency,
		&user.CreatedAt,
		&nullData.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) user: %v", identifier, err)
		return
	}

	// Check if isn't empty to access its value
	if nullData.AvatarURL != nil {
		user.Avatar.URL = nullData.AvatarURL.String
	}
	if nullData.Name != nil {
		user.Name = nullData.Name.String
	}
	if nullData.LastName != nil {
		user.LastName = nullData.LastName.String
	}
	if nullData.Birthday != nil {
		user.Birthday = &nullData.Birthday.Time
	}
	if nullData.UpdatedAt != nil {
		user.UpdatedAt = &nullData.UpdatedAt.Time
	}
	return
}
