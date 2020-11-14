package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/lib/pq"
)

// SelectProduct - Selects a product.
func SelectProduct(db *sql.DB, productToFind models.Product) (product models.Product, err error) {
	if db == nil {
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
			id, user_id, name, description, categories, price, images, created_at, updated_at
		FROM
			products
		WHERE
			id = $1
	`

	stmt, err := db.Prepare(query)
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
		&product.Description,
		pq.Array(&product.Categories),
		&product.Price,
		pq.Array(&product.Images),
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
