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

	// This query sets the database fields to its last value if
	// the param is empty. Otherwise, sets the param value.
	query := fmt.Sprintf(`
		UPDATE
			users
		SET
			language = CASE WHEN $1 = '' THEN language ELSE $1 END,
			avatar = CASE WHEN $2 = '' THEN avatar ELSE $2 END,
			username = CASE WHEN $3 = '' THEN username ELSE $3 END,
			password = CASE WHEN $4 = '' THEN password ELSE $4 END,
			email = CASE WHEN $5 = '' THEN email ELSE $5 END,
			name = CASE WHEN $6 = '' THEN name ELSE $6 END,
			last_name = CASE WHEN $7 = '' THEN last_name ELSE $7 END,
			birthday = CASE WHEN $8::timestamp IS NULL THEN birthday ELSE $8 END,
			theme = CASE WHEN $9 = '' THEN theme ELSE $9 END,
			currency = CASE WHEN $10 = '' THEN currency ELSE $10 END,
			updated_at = NOW()
		WHERE 
			id = $11 OR username = $12 OR email = $13
		RETURNING
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

// Helper function that sets the query values to put on the query,
// if a field isn't empty, it use the user value.
// Otherwise, use the field database name for don't changes.
func getUpdateUserParams(user models.User) (values map[string]interface{}) {
	values = make(map[string]interface{})

	if user.Language != "" {
		values["language"] = user.Language
	} else {
		values["language"] = "language"
	}

	if user.Avatar.URL != "" {
		values["avatar"] = user.Avatar.URL
	} else {
		values["avatar"] = "avatar"
	}

	if user.Username != "" {
		values["username"] = user.Username
	} else {
		values["username"] = "username"
	}

	if user.Password != "" {
		values["password"] = user.Password
	} else {
		values["password"] = "password"
	}

	if user.Email != "" {
		values["email"] = user.Email
	} else {
		values["email"] = "email"
	}

	if user.Birthday != nil {
		values["birthday"] = user.Birthday
	} else {
		values["birthday"] = "birthday"
	}

	return
}
