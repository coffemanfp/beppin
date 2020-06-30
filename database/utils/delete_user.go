package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteUser - Deletes a user.
func DeleteUser(db *sql.DB, userID int, username string) (err error) {
	exists, err := ExistsUser(db, userID, username)
	if err != nil {
		return
	}

	if !exists {
		err = errors.New(errs.ErrNotExistentObject)
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
		err = fmt.Errorf("failed to prepare the delete user statement:\n%s", err)

		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, username)
	if err != nil {
		err = fmt.Errorf("failed to delete the user:\n%s", err)
	}
	return
}
