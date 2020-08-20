package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
)

// DeleteProduct - Deletes a product.
func DeleteProduct(db *sql.DB, product models.Product) (err error) {
	identifier := product.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to delete product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		DELETE FROM
			products
		WHERE
			id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(product.ID)
	if err != nil {
		err = fmt.Errorf("failed to delete (%v) product: %v", identifier, err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get the rows affected number: %v", err)
		return
	}

	if rowsAffected == 0 {
		err = fmt.Errorf("failed to delete (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
	}
	return
}
