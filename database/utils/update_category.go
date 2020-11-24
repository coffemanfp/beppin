package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateCategory - Updates a category.
func UpdateCategory(db *sql.DB, categoryToUpdate, category models.Category) (updatedCategory models.Category, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := categoryToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update category: %w (category)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	// This query sets the database fields to its last value if
	// the param is empty. Otherwise, sets the param value.
	query := `
		UPDATE
			categories
		SET
			name = CASE WHEN $1 = '' THEN name ELSE $1 END,
			description = CASE WHEN $2 = '' THEN description ELSE $2 END,
			updated_at = NOW()
		WHERE 
			id =  $3
		RETURNING
			id, path, created_at, updated_at
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) category statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullCategoryData

	err = stmt.QueryRow(
		category.Name,
		category.Description,
		categoryToUpdate.ID,
	).Scan(
		&updatedCategory.ID,
		&updatedCategory.Name,
		&nullData.Description,
		&updatedCategory.CreatedAt,
		&updatedCategory.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) category: %w (category)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) category: %v", identifier, err)
		return
	}
	nullData.setResults(&updatedCategory)
	return
}
