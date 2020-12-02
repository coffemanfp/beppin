package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateUser - Updates a user.
func UpdateUser(dbtx DBTX, userToUpdate, user models.User) (userUpdated models.User, err error) {
	if dbtx == nil {
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
		WITH updated AS (
		UPDATE
			users
		SET
			language = CASE WHEN $1 = '' THEN language ELSE $1 END,
			avatar_id = CASE WHEN $2 = 0 THEN avatar_id ELSE $2 END,
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
			id, language, avatar_id, username, email, name, last_name, birthday, theme, currency, created_at, updated_at
		)
		SELECT
			updated.*, files.path
		FROM
			updated
		LEFT JOIN
			files
		ON
			updated.avatar_id = files.id
	`)

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullUserData

	if user.Avatar == nil {
		user.Avatar = new(models.File)
	}

	err = stmt.QueryRow(
		user.Language,
		user.Avatar.ID,
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
	).Scan(
		&userUpdated.ID,
		&userUpdated.Language,
		&nullData.AvatarID,
		&userUpdated.Username,
		&userUpdated.Email,
		&nullData.Name,
		&nullData.LastName,
		&nullData.Birthday,
		&userUpdated.Theme,
		&userUpdated.Currency,
		&userUpdated.CreatedAt,
		&userUpdated.UpdatedAt,
		&nullData.AvatarPath,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) user: %v", identifier, err)
		return
	}

	nullData.setResults(&userUpdated)
	if userUpdated.Avatar != nil {
		userUpdated.Avatar.SetURL()
	}
	return
}
