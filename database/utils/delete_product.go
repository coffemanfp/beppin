package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteProduct - Deletess a product.
func DeleteProduct(db *sql.DB, productID int) (err error) {
	exists, err := ExistsProduct(db, productID)
	if err != nil {
		return
	}

	if !exists {
		err = errors.New(errs.ErrNotExistentObject)
		return
	}

	query := `
		DELETE FROM
			products
		WHERE
			products.id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete product statement:\n%s", err)

		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(productID)
	if err != nil {
		err = fmt.Errorf("failed to delete the product:\n%s", err)
	}
	return
}
