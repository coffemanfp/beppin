package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateLanguage - Updates a language.
func UpdateLanguage(dbtx DBTX, languageToUpdate, language models.Language) (updatedLanguage models.Language, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := languageToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	// This query sets the database fields to its last value if
	// the param is empty. Otherwise, sets the param value.
	query := `
		UPDATE
			languages
		SET
			code = CASE WHEN $1 = '' THEN code ELSE $1 END,
			status = CASE WHEN $2 = '' THEN status ELSE $2 END,
			updated_at = NOW()
		WHERE 
			id =  $3
		RETURNING
			id, code, status, created_at, updated_at
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		language.Code,
		language.Status,
		languageToUpdate.ID,
	).Scan(
		&updatedLanguage.ID,
		&updatedLanguage.Code,
		&updatedLanguage.Status,
		&updatedLanguage.CreatedAt,
		&updatedLanguage.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) language: %w (language)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) language: %v", identifier, err)
	}

	return
}
