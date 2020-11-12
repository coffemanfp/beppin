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
			id, language, avatar, username, email, name, last_name, birthday, theme, currency, created_at, updated_at FROM users
		WHERE
			id = $1 OR username = $2 OR email = $3
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullUserData

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

	nullData.setResults(&user)
	return
}
