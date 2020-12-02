package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// InsertProductCategory - Inserts a product category.
func InsertProductCategory(dbtx DBTX, productID int64, categoryID int64) (err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			product_categories(category_id, product_id)
		VALUES
			($1, $2)
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert product_category statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryID, productID)
	if err != nil {
		err = fmt.Errorf("failed to execute insert product_category statement: %v", err)
	}
	return
}
