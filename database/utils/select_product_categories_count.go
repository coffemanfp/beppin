package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// SelectProductCategoriesCount - Selects the product categories relations count.
func SelectProductCategoriesCount(dbtx DBTX, productID int64) (count int, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	if productID == 0 {
		err = fmt.Errorf("failed to check product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
				SELECT
					COUNT(*)
				FROM
					product_categories
				WHERE
					product_id = $1
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists product_category statement: %v", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(productID).Scan(&count)
	if err != nil {
		err = fmt.Errorf("failed to select the exists product_categories statement: %v", err)
	}
	return
}
