package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectProductCategories - Selects the categories of a product.
func SelectProductCategories(dbtx DBTX, productToFind models.Product) (categories models.Categories, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := productToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			categories.id, categories.name
		FROM
			categories
		INNER JOIN
			products
		INNER JOIN
			categories_products
		ON
			categories_products.product_id = products.id
		ON
			categories_products.category_id = categories.id
		WHERE
			products.id = $1
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) product statement: %v", identifier, err)

		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(productToFind.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) product: %v", identifier, err)
	}

	var category models.Category

	for rows.Next() {
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan category: %v", err)
			return
		}

		categories = append(categories, category)

		// Empty the value to avoid overwrite
		category = models.Category{}
	}
	return
}
