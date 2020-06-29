package utils

import (
	"database/sql"
	"fmt"
)

// ExistsProduct - Checks if exists a product.
func ExistsProduct(db *sql.DB, productID int) (exists bool, err error) {
	query := `
		SELECT
			EXISTS(
				SELECT
					id
				FROM
					products
				WHERE
					id = $1
			)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the exists product statement:\n%s", err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(productID).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("failed to select the exists product statement:\n%s", err)
	}
	return
}

