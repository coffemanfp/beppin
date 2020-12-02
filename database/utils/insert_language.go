package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// InsertLanguage - Insert a language.
func InsertLanguage(dbtx DBTX, language models.Language) (createdLanguage models.Language, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := language.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to insert language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		INSERT INTO
			languages(code, status)
		VALUES
			($1, $2)
		ON CONFLICT DO
			NOTHING
		RETURNING
			id, code, status, created_at
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		language.Code,
		language.Status,
	).Scan(
		&createdLanguage.ID,
		&createdLanguage.Code,
		&createdLanguage.Status,
		&createdLanguage.CreatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert (%v) language statement: %v", identifier, err)
	}
	return
}
