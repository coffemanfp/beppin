package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
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
			id, code, status, created_at, updated_at
		FROM
			languages
		WHERE
			id = $1 OR code = $1
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullLanguageData

	err = stmt.QueryRow(
		languageToFind.ID,
		languageToFind.Code,
	).Scan(
		&language.ID,
		&language.Code,
		&language.Status,
		&language.CreatedAt,
		&nullData.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) language: %w (language)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) language: %v", identifier, err)
	}

	nullData.setResults(&language)
	return
}
