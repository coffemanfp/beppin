package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteProduct - Deletess a product.
func DeleteProduct(db *sql.DB, productID int) (err error) {
	query := `
		DELETE FROM
			products
		WHERE
			id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete product statement:\n%s", err)

		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(productID)
	if err != nil {
		err = fmt.Errorf("failed to delete the product:\n%s", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get the rows affected number:\n%s", err)
		return
	}

	if rowsAffected == 0 {
		err = errors.New(errs.ErrNotExistentObject)
	}
	return
}
