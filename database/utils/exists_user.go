package utils

import (
	"database/sql"
	"fmt"
)

// ExistsUser - Checks if exists a user.
func ExistsUser(db *sql.DB, userID int, username string) (exists bool, err error) {
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
		err = fmt.Errorf("failed to prepare the exists user statement:\n%s", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID, username).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists user statement:\n%s", err)
	}
	return
}
