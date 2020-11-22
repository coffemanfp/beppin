package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateFile - Updates a file.
func UpdateFile(db *sql.DB, fileToUpdate, file models.File) (updatedFile models.File, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := fileToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update file: %w (file)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	// This query sets the database fields to its last value if
	// the param is empty. Otherwise, sets the param value.
	query := `
		UPDATE
			files
		SET
			path = CASE WHEN $1 = '' THEN path ELSE $1 END,
			updated_at = NOW()
		WHERE 
			id =  $2
		RETURNING
			id, path, created_at, updated_at
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) file statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullFileData

	err = stmt.QueryRow(
		file.Path,
		fileToUpdate.ID,
	).Scan(
		&updatedFile.ID,
		&updatedFile.Path,
		&updatedFile.CreatedAt,
		&updatedFile.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) file: %w (file)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) file: %v", identifier, err)
		return
	}
	nullData.setResults(&updatedFile)
	updatedFile.SetURL()
	return
}
