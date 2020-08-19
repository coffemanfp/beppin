package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// UpdateAvatar updates the avatar url.
func UpdateAvatar(db *sql.DB, avatarURL string, userToUpdate models.User) (err error) {
	identifier := userToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update user avatar: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		UPDATE
			users
		SET
			avatar = $1
		WHERE
			id = $2 OR username = $3 OR email = $4
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user avatar statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		avatarURL,
		userToUpdate.ID,
		userToUpdate.Username,
		userToUpdate.Email,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute the update (%v) user avatar statement: %v", identifier, err)
	}
	return
}
