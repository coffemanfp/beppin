package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/models"
	"github.com/lib/pq"
)

// SelectProduct - Selects a product.
func SelectProduct(db *sql.DB, productID int) (product models.Product, err error) {
	query := `
		SELECT
			user_id, name, description, categories
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
		&product.Name,
		&product.Description,
		(*pq.StringArray)(&product.Categories),
	)
	if err != nil {
		err = fmt.Errorf("failed to select the product:\n%s", err)
	}
	return
}
