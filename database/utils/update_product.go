package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/lib/pq"
)

// UpdateProduct - Updates a product.
func UpdateProduct(db *sql.DB, productToUpdate, product models.Product) (id int, err error) {
	if db == nil {
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
			categories = CASE WHEN $3::varchar[] IS NULL THEN categories ELSE $3 END,
			price = CASE WHEN $4 = 0.0 THEN price ELSE $4 END,
			updated_at = NOW()
		WHERE 
			id =  $5
		RETURNING
			id
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		product.Name,
		product.Description,
		pq.Array(product.Categories),
		product.Price,
		productToUpdate.ID,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to update (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to update (%v) product: %v", identifier, err)
		return
	}
	return
}
