package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteUser - Deletes a user.
func DeleteUser(db *sql.DB, userID int, username string) (err error) {
	query := `
		DELETE FROM
			users
		WHERE
			users.id = $1 OR users.username = $2
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete user statement:\n%s", err)

		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(userID, username)
	if err != nil {
		err = fmt.Errorf("failed to delete the user:\n%s", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get the rows affected number:\n%s", err)
		return
	}

	if rowsAffected == 0 {
		err = errs.ErrNotExistentObject
	}
	return
}
