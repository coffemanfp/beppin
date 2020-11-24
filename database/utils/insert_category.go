package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// InsertCategory - Inserts a category.
func InsertCategory(db *sql.DB, category models.Category) (createdCategory models.Category, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
			INSERT INTO
					categories(name, description)
			VALUES
					($1)
			RETURNING
					id, name, description, created_at
    `

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert category statement: %v", err)
		return
	}
	defer stmt.Close()

	var nullData nullCategoryData

	err = stmt.QueryRow(
		category.Name,
	).Scan(
		&createdCategory.ID,
		&createdCategory.Name,
		&nullData.Description,
		&createdCategory.CreatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert category statement: %v", err)
	}

	nullData.setResults(&createdCategory)
	return
}
