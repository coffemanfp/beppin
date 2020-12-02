package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// DeleteProductCategories - Deletes all the categories related with a product.
func DeleteProductCategories(dbtx DBTX, productID int64) (err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		DELETE FROM
			product_categories
		WHERE
			product_id = $1
		RETURNING
			id
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) product_categories statement: %v", productID, err)
		return
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRow(productID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) product_categories: %w", productID, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) product_categories: %v", productID, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) product_categories: %w", productID, errs.ErrNotExistentObject)
	}
	return
}
