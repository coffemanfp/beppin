package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectCategory - Selects a category.
func SelectCategory(dbtx DBTX, categoryToFind models.Category) (category models.Category, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := categoryToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select category: %w (category)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, name, description, created_at, updated_at
		FROM
			categories
		WHERE
			id = $1 OR path = $2
			
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) category statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullCategoryData

	err = stmt.QueryRow(
		categoryToFind.ID,
		categoryToFind.Name,
	).Scan(
		&category.ID,
		&category.Name,
		&nullData.Description,
		&category.CreatedAt,
		&nullData.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) category: %w (category)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) category: %v", identifier, err)
	}

	nullData.setResults(&category)
	return
}
