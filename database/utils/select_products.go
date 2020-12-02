package utils

import (
	"database/sql"
	"errors"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// SelectProducts - Select a products list.
func SelectProducts(dbtx DBTX, limit, offset int) (products models.Products, err error) {
	if dbtx == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
	SELECT
		id, user_id, name, description, price, created_at, updated_at
	FROM
		products
	ORDER BY
		id
	LIMIT
		$1
	OFFSET
		$2
	`

	stmt, err := dbtx.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select products statement: %v", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("failed to select products: %w", errs.ErrNotExistentObject)
			return
		}

		err = fmt.Errorf("failed to select products: %v", err)
		return
	}

	var product models.Product
	var nullData nullProductData

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.UserID,
			&product.Name,
			&nullData.Description,
			&product.Price,
			&product.CreatedAt,
			&nullData.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan product: %v", err)
			return
		}

		nullData.setResults(&product)
		products = append(products, product)

		// Empty the value to avoid overwrite
		product = models.Product{}
	}

	return
}
