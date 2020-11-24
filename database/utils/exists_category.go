package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// ExistsCategory - Checks if exists a category.
func ExistsCategory(db *sql.DB, category models.Category) (exists bool, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := category.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check category: %w (category)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					categories
				WHERE
					id = $1 OR name = $2
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) category statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		category.ID,
		category.Name,
	).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) category statement: %v", identifier, err)
	}
	return
}
