package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// ExistsLanguage - Checks if exists a language.
func ExistsLanguage(dbtx DBTX, language models.Language) (exists bool, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := language.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check language: %w (language)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					languages
				WHERE
					code = $1
			)
	`

	stmt, err := dbtx.Prepare(query)
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
