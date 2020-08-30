package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// SelectLanguage - Selects a language.
func SelectLanguage(db *sql.DB, languageToFind models.Language) (language models.Language, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := languageToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

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
		err = fmt.Errorf("failed to prepare the select (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(languageToFind.Code).Scan(
		&language.Code,
		&language.Status,
		&language.CreatedAt,
		&language.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) language: %w (language)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) language: %v", identifier, err)
	}
	return
}
