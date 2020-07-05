package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// SelectLanguage - Selects a language.
func SelectLanguage(db *sql.DB, languageCode string) (language models.Language, err error) {
	query := `
		SELECT
			code, status, created_at, updated_at
		FROM
			languages
		WHERE
			code = $1
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select language statement:\n%s", err)

		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(languageCode).Scan(
		&language.Code,
		&language.Status,
		&language.CreatedAt,
		&language.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New(errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select the language:\n%s", err)
	}
	return
}
