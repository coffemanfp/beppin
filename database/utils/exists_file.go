package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// ExistsFile - Checks if exists a file.
func ExistsFile(db *sql.DB, file models.File) (exists bool, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := file.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check file: %w (file)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					files
				WHERE
					id = $1 OR path = $2
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) file statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		file.ID,
		file.Path,
	).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) file statement: %v", identifier, err)
	}
	return
}
