package utils

import (
	"database/sql"
	"fmt"
)

// TruncateLanguages deletes all the language records.
func TruncateLanguages(db *sql.DB, cascade bool) (err error) {
	var query string

	if cascade {
		query = `
			TRUNCATE TABLE
				languages
			CASCADE
		`
	} else {
		query = `
			TRUNCATE TABLE
				languages
		`
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the truncate table languages statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		err = fmt.Errorf("failed to exec the truncate table languages statement: %v", err)
	}
	return
}
