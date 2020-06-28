package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/models"
)

// SelectProducts - Select a products list.
func SelectProducts(db *sql.DB, limit int, offset int) (products models.Products, err error) {
	query := `
		SELECT
			id, user_id, name, description, created_at, updated_at
		FROM
			products
		LIMIT
			$1
		OFFSET
			$2
	`

	settings, err := config.GetSettings()
	if err != nil {
		return
	}

	if limit == 0 {
		limit = settings.MaxElementsPerPagination
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the select products statement:\n%s", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		err = fmt.Errorf("failed to select the products:\n%s", err)
		return
	}

	var product models.Product

	for rows.Next() {
		err = rows.Scan(
			product.ID,
			product.UserID,
			product.Name,
			product.Description,
			product.CreatedAt,
			product.UpdatedAt,
		)
		if err != nil {
			err = fmt.Errorf("failed to scan a product:\n%s", err)
			return
		}

		products = append(products, product)
	}

	return
}
