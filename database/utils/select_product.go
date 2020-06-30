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
func SelectProduct(db *sql.DB, productID int) (product models.Product, err error) {
	exists, err := ExistsProduct(db, productID)
	if err != nil {
		return
	}

	if !exists {
		err = errors.New(errs.ErrNotExistentObject)
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
		err = fmt.Errorf("failed to prepare the select product statement:\n%s", err)

		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(productID).Scan(
		&product.ID,
		&product.UserID,
		&product.Name,
		&product.Description,
		(*pq.StringArray)(&product.Categories),
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to select the product:\n%s", err)
	}
	return
}
