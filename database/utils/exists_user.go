package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// ExistsUser - Checks if exists a user.
func ExistsUser(db *sql.DB, user models.User) (exists bool, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					users
				WHERE
					id = $1 OR username = $2 OR email = $3
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		user.ID,
		user.Username,
		user.Email,
	).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) user statement: %v", identifier, err)
	}
	return
}
