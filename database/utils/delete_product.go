package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/errors"
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
		err = fmt.Errorf("failed to prepare the delete %d product statement: %v", productID, err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(productID)
	if err != nil {
		err = fmt.Errorf("failed to delete %d product: %v", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get the rows affected number:\n%v", err)
		return
	}

	if rowsAffected == 0 {
		err = fmt.Errorf("failed to delete %d product: %w", productID, errors.ErrNotExistentObject)
	}
	return
}
