package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// DeleteLanguage - Deletes a language.
func DeleteLanguage(db *sql.DB, language models.Language) (id int, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := language.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to delete language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		DELETE FROM
			languages
		WHERE
			languages.id = $1 OR languages.code = $2
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		language.ID,
		language.Code,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) language: %w (language)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) language: %v", identifier, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) language: %w (language)", identifier, errs.ErrNotExistentObject)
	}
	return
}
