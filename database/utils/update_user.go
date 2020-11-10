package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateUser - Updates a user.
func UpdateUser(db *sql.DB, userToUpdate, user models.User) (userUpdated models.User, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := userToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			language = $1,
			avatar = $2,
			username = $3,
			password = $4,
			email = $5,
			name = $6,
			last_name = $7,
			birthday = $8,
			theme = $9,
			currency = $10
			updated_at = NOW()
		WHERE 
			id = $11 OR username = $12 OR email = $13
		RETUNING
			id
	`)

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRow(
		user.Language,
		user.Avatar.URL,
		user.Username,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Birthday,
		user.Theme,
		user.Currency,
		userToUpdate.ID,
		userToUpdate.Username,
		userToUpdate.Email,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) user: %v", identifier, err)
		return
	}

	userUpdated = user
	userUpdated.ID = id
	return
}
