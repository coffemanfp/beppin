package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/lib/pq"
)

// SelectProducts - Select a products list.
func SelectProducts(db *sql.DB, limit int, offset int) (products models.Products, err error) {
	query := `
		SELECT
			id, user_id, name, description, categories, created_at, updated_at
		FROM
			products
		LIMIT
			$1
		OFFSET
			$2
	`

	settings := config.GetSettings()

	if limit == 0 {
		limit = settings.MaxElementsPerPagination
	}

	stmt, err := db.Prepare(query)
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

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.UserID,
			&product.Name,
			&product.Description,
			pq.Array(&product.Categories),
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan product: %v", err)
			return
		}

		products = append(products, product)
	}

	return
}
