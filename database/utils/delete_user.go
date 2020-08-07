package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteUser - Deletes a user.
func DeleteUser(db *sql.DB, user models.User) (err error) {
	identifier := user.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to delete user: %w (user)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		DELETE FROM
			users
		WHERE
			users.id = $1 OR users.username = $2
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) user statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.ID, user.Username)
	if err != nil {
		err = fmt.Errorf("failed to delete (%v) user: %v", identifier, err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get the rows affected number: %v", err)
		return
	}

	if rowsAffected == 0 {
		err = fmt.Errorf("failed to delete (%v) user: %w (user)", identifier, errs.ErrNotExistentObject)
	}
	return
}
