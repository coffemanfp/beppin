package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
)

// SelectLanguage - Selects a language.
func SelectLanguage(db *sql.DB, languageID int, languageCode string) (language models.Language, err error) {
	query := `
		SELECT
			id, code, status, created_at, updated_at
		FROM
			languages
		WHERE
			id = $1
			OR
			code = $2
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select language statement:\n%s", err)

		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(languageID, languageCode).Scan(
		&language.ID,
		&language.Code,
		&language.Status,
		&language.CreatedAt,
		&language.UpdatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to select the language:\n%s", err)
	}
	return
}
