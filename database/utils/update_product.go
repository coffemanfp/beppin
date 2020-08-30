package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
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

	query := `
		UPDATE
			products
		SET
			name = $1,
			description = $2,
			categories = $3,
			updated_at = NOW()
		WHERE 
			id =  $4
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
		productToUpdate.ID,
	).Scan(&id)
	if err != nil {
		err = fmt.Errorf("failed to execute the update (%v) product statement: %v", identifier, err)
	}
	return
}
