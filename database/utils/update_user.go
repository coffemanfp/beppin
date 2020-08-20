package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// UpdateUser - Updates a user.
func UpdateUser(db *sql.DB, userToUpdate, user models.User) (err error) {
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
			updated_at = NOW()
		WHERE 
			id = $10 OR username = $11 OR email = $12
	`)

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Language.Code,
		user.AvatarURL,
		user.Username,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Birthday,
		user.Theme,
		userToUpdate.ID,
		userToUpdate.Username,
		userToUpdate.Email,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute the update (%v) user statement: %v", identifier, err)
	}
	return
}
