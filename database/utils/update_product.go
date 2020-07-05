package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	"github.com/lib/pq"
)

// UpdateProduct - Updates a product.
func UpdateProduct(db *sql.DB, productID int, product models.Product) (err error) {
	previosProductData, err := SelectProduct(db, productID)
	if err != nil {
		return
	}

	product = fillProductEmptyFields(product, previosProductData)

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
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("failed to prepare the update product statement:\n%s", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.Name,
		product.Description,
		pq.Array(product.Categories),
		productID,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute the update product statement:\n%s", err)
	}
	return
}

func fillProductEmptyFields(product models.Product, previousProductData models.Product) models.Product {

	switch "" {
	case product.Name:
		product.Name = previousProductData.Name
	case product.Description:
		product.Description = previousProductData.Description
	}

	if product.Categories == nil {
		product.Categories = previousProductData.Categories
	}

	return product
}
