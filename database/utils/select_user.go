package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// SelectUser - Selects a user.
func SelectUser(db *sql.DB, userToFind models.User) (user models.User, err error) {
	identifier := userToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, language, avatar, username, email, name, last_name, birthday, theme, created_at, updated_at
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

	err = stmt.QueryRow(
		userToFind.ID,
		userToFind.Username,
		userToFind.Email,
	).Scan(
		&user.ID,
		&user.Language.Code,
		&user.AvatarURL,
		&user.Username,
		&user.Email,
		&user.Name,
		&user.LastName,
		&user.Birthday,
		&user.Theme,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) user: %v", identifier, err)
		return
	}
	return
}
