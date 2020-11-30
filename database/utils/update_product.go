package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// UpdateProduct - Updates a product.
func UpdateProduct(dbtx DBTX, productToUpdate, product models.Product) (updatedProduct models.Product, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	identifier := productToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	// This query sets the database fields to its last value if
	// the param is empty. Otherwise, sets the param value.
	query := `
		UPDATE
			products
		SET
			name = CASE WHEN $1 = '' THEN name ELSE $1 END,
			description = CASE WHEN $2 = '' THEN description ELSE $2 END,
			price = CASE WHEN $3 = 0.0 THEN price ELSE $3 END,
			updated_at = NOW()
		WHERE 
			id =  $4
		RETURNING
			id, user_id, name, description, price, created_at, updated_at
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	var nullData nullProductData

	err = stmt.QueryRow(
		product.Name,
		product.Description,
		product.Price,
		productToUpdate.ID,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.UserID,
		&updatedProduct.Name,
		&nullData.Description,
		&updatedProduct.Price,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) product: %v", identifier, err)
		return
	}
	nullData.setResults(&updatedProduct)
	return
}
