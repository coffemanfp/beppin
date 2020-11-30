package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectCategories - Select a categories list.
func SelectCategories(dbtx DBTX) (categories models.Categories, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
	SELECT
		id, name, description, created_at, updated_at
	FROM
		categories
	ORDER BY
		id
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select categories statement: %v", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select categories: %w", errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select categories: %v", err)
		return
	}

	var category models.Category
	var nullData nullCategoryData

	for rows.Next() {
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&nullData.Description,
			&category.CreatedAt,
			&nullData.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan category: %v", err)
			return
		}

		nullData.setResults(&category)
		categories = append(categories, category)

		// Empty the value to avoid overwrite
		category = models.Category{}
	}

	return
}
