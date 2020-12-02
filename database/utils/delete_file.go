package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// DeleteFile - Deletes a file.
func DeleteFile(dbtx DBTX, file models.File) (id int, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := file.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to delete file: %w (file)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		DELETE FROM
			files
		WHERE
			files.id = $1 OR files.path = $2
		RETURNING
			id
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) file statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		file.ID,
		file.Path,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) file: %w (file)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) file: %v", identifier, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) file: %w (file)", identifier, errs.ErrNotExistentObject)
	}
	return
}
