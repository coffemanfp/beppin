package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
)

// InsertProductCategory - Inserts a product category.
func InsertProductCategory(db *sql.DB, productID int64, categoryID int64) (err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			categories_products(category_id, product_id)
		VALUES
			($1, $2)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert category product statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryID, productID)
	if err != nil {
		err = fmt.Errorf("failed to execute insert category product statement: %v", err)
	}
	return
}
