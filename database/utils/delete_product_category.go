package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// DeleteProductCategory - Deletes a category related with a product.
func DeleteProductCategory(dbtx DBTX, productID int64, categoryID int64) (err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		DELETE FROM
			product_categories
		WHERE
			product_id = $1 AND category_id = $2
		RETURNING
			id
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) product_category statement: %v", productID, err)
		return
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRow(productID, categoryID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) product_category: %w", productID, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) product_category: %v", productID, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) product_category: %w (product_category)", productID, errs.ErrNotExistentObject)
	}
	return
}
