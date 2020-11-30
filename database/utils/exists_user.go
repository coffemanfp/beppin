package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// ExistsUser - Checks if exists a user.
func ExistsUser(dbtx DBTX, user models.User) (exists bool, err error) {
	if dbtx == nil {
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

	stmt, err := dbtx.Prepare(query)
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
