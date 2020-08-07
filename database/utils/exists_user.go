package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// ExistsUser - Checks if exists a user.
func ExistsUser(db *sql.DB, user models.User) (exists bool, err error) {
	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					id
				FROM
					users
				WHERE
					id = $1 OR username = $2
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.ID, user.Username).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) user statement: %v", identifier, err)
	}
	return
}
