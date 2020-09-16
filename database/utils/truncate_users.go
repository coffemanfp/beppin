package utils

import (
	"database/sql"
	"fmt"
)

// TruncateUsers deletes all the user records.
func TruncateUsers(db *sql.DB, cascade bool) (err error) {
	var query string

	if cascade {
		query = `
			TRUNCATE TABLE
				users
			CASCADE
		`
	} else {
		query = `
			TRUNCATE TABLE
				users
		`
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the truncate table users statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		err = fmt.Errorf("failed to exec the truncate table users statement: %v", err)
	}
	return
}
