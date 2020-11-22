package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectFile - Selects a file.
func SelectFile(db *sql.DB, fileToFind models.File) (file models.File, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := fileToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select file: %w (file)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, path, created_at, updated_at
		FROM
			files
		WHERE
			id = $1 OR path = $2
			
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) file statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullFileData

	err = stmt.QueryRow(
		fileToFind.ID,
		fileToFind.Path,
	).Scan(
		&file.ID,
		&file.Path,
		&file.CreatedAt,
		&nullData.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) file: %w (file)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) file: %v", identifier, err)
	}

	nullData.setResults(&file)
	file.SetURL()
	return
}
