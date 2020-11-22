package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// InsertFile - Inserts a file.
func InsertFile(db *sql.DB, file models.File) (createdFile models.File, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
			INSERT INTO
					files(path)
			VALUES
					($1)
			RETURNING
					id, path, created_at
    `

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert file statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		file.Path,
	).Scan(
		&createdFile.ID,
		&createdFile.Path,
		&createdFile.CreatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert file statement: %v", err)
	}
	createdFile.SetURL()
	return
}
