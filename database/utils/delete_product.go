package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
)

// DeleteProduct - Deletes a product.
func DeleteProduct(db *sql.DB, product models.Product) (id int, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

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
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the delete (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(product.ID).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to delete (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to delete (%v) product: %v", identifier, err)
		return
	}

	if id == 0 {
		err = fmt.Errorf("failed to delete (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
	}
	return
}
