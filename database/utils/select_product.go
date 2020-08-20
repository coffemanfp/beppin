package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/lib/pq"
)

// SelectProduct - Selects a product.
func SelectProduct(db *sql.DB, productToFind models.Product) (product models.Product, err error) {
	identifier := productToFind.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to select product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	query := `
		SELECT
			id, user_id, name, description, categories, created_at, updated_at
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

	err = stmt.QueryRow(productToFind.ID).Scan(
		&product.ID,
		&product.UserID,
		&product.Name,
		&product.Description,
		(*pq.StringArray)(&product.Categories),
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select (%v) product: %w (product)", identifier, errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select (%v) product: %v", identifier, err)
	}
	return
}
