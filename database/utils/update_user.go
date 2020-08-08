package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// UpdateUser - Updates a user.
func UpdateUser(db *sql.DB, userToUpdate, user models.User) (err error) {
	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	previousUserData, err := SelectUser(db, userToUpdate)
	if err != nil {
		return
	}

	user = fillUserEmptyFields(user, previousUserData)

	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			language = $1,
			username = $2,
			password = $3,
			email = $4,
			name = $5,
			last_name = $6,
			birthday = $7,
			theme = $8,
			updated_at = NOW()
		WHERE 
			id = $9 OR username = $10 OR email = $11
	`)

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.Language.Code,
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

func fillUserEmptyFields(user, previousUserData models.User) models.User {

	switch "" {
	case user.Language.Code:
		user.Language.Code = previousUserData.Language.Code

	case user.Username:
		user.Username = previousUserData.Username

	case user.Email:
		user.Email = previousUserData.Email

	case user.Password:
		user.Password = previousUserData.Password

	case user.Name:
		user.Name = previousUserData.Name

	case user.LastName:
		user.LastName = previousUserData.LastName

	case user.Theme:
		user.Theme = previousUserData.Theme
	}

	if user.Birthday == nil {
		user.Birthday = previousUserData.Birthday
	}

	return user
}
