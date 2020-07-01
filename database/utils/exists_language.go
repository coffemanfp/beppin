package utils

import (
	"database/sql"
	"fmt"
)

// ExistsLanguage - Checks if exists a language.
func ExistsLanguage(db *sql.DB, languageID int, languageCode string) (exists bool, err error) {
	query := `
		SELECT
			EXISTS(
				SELECT
					id
				FROM
					languages
				WHERE
					id = $1 OR language_code = $2
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists language statement:\n%s", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(languageID, languageCode).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists language statement:\n%s", err)
	}
	return
}
