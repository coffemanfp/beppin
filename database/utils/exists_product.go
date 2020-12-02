package utils

import (
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// ExistsProduct - Checks if exists a product.
func ExistsProduct(dbtx DBTX, product models.Product) (exists bool, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := product.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to check product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			EXISTS(
				SELECT
					1
				FROM
					products
				WHERE
					id = $1
			)
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(product.ID).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists (%v) product statement: %v", identifier, err)
	}
	return
}
