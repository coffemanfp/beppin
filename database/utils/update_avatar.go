package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateAvatar updates the avatar url.
func UpdateAvatar(db *sql.DB, avatarURL string, userToUpdate models.User) (id int, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := userToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update user avatar: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		UPDATE
			users
		SET
			avatar = $1,
			updated_at = NOW()
		WHERE
			id = $2 OR username = $3 OR email = $4
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user avatar statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		avatarURL,
		userToUpdate.ID,
		userToUpdate.Username,
		userToUpdate.Email,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) avatar user: %w (user)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) avatar user: %v", identifier, err)
		return
	}
	return
}
