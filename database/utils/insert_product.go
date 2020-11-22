package utils

import (
	"database/sql"
	"fmt"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

// InsertProduct - Insert a product.
func InsertProduct(db *sql.DB, product models.Product) (createdProduct models.Product, err error) {
	if db == nil {
		err = errs.ErrClosedDatabase
		return
	}

	query := `
		INSERT INTO
			products(user_id, name, description, price)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, user_id, name, description, price, created_at
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the insert product statement: %v", err)
		return
	}
	defer stmt.Close()

	var nullData nullProductData

	err = stmt.QueryRow(
		product.UserID,
		product.Name,
		product.Description,
		product.Price,
	).Scan(
		&createdProduct.ID,
		&createdProduct.UserID,
		&createdProduct.Name,
		&nullData.Description,
		&createdProduct.Price,
		&createdProduct.CreatedAt,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute insert product statement: %v", err)
	}

	nullData.setResults(&createdProduct)
	return
}
