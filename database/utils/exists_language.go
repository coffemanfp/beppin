package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// ExistsLanguage - Checks if exists a language.
func ExistsLanguage(db *sql.DB, language models.Language) (exists bool, err error) {
	identifier := language.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					id
				FROM
					languages
				WHERE
					language_code = $2
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) language statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(language.Code).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) language statement: %v", identifier, err)
	}
	return
}
