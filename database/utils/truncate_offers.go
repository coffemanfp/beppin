package utils

import (
	"database/sql"
	"fmt"
)

// TruncateOffers deletes all the offer records.
func TruncateOffers(db *sql.DB, cascade bool) (err error) {
	var query string

	if cascade {
		query = `
			TRUNCATE TABLE
				offers
			CASCADE
		`
	} else {
		query = `
			TRUNCATE TABLE
				offers
		`
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the truncate table offers statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		err = fmt.Errorf("failed to exec the truncate table offers statement: %v", err)
	}
	return
}
