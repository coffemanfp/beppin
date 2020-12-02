package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// ExistsProductCategory - Checks if a product have a category.
func ExistsProductCategory(dbtx DBTX, productID int64, categoryID int64) (exists bool, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	if productID == 0 {
		err = fmt.Errorf("failed to check product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}
	if categoryID == 0 {
		err = fmt.Errorf("failed to check category: %w (category)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					product_categories
				WHERE
					product_id = $1 AND category_id = $2
			)
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists product_category statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		productID,
		categoryID,
	).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists product_category statement: %v", err)
	}
	return
}
