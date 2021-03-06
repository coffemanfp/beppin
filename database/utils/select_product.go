package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectProduct - Selects a product.
func SelectProduct(dbtx DBTX, productToFind models.Product) (product models.Product, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := productToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, user_id, name, description, price, created_at, updated_at
		FROM
			products
		WHERE
			id = $1
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select (%v) product statement: %v", identifier, err)

		return
	}
	defer stmt.Close()

	var nullData nullProductData

	err = stmt.QueryRow(productToFind.ID).Scan(
		&product.ID,
		&product.UserID,
		&product.Name,
		&nullData.Description,
		&product.Price,
		&product.CreatedAt,
		&nullData.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) product: %v", identifier, err)
	}

	nullData.setResults(&product)
	return
}
