package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// DeleteCategory - Deletes a category.
func DeleteCategory(db *sql.DB, category models.Category) (id int, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := category.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to delete category: %w (category)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		DELETE FROM
			categories
		WHERE
			categorys.id = $1 OR categorys.name = $2
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) category statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		category.ID,
		category.Name,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) category: %w (category)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) category: %v", identifier, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) category: %w (category)", identifier, errs.ErrNotExistentObject)
	}
	return
}
