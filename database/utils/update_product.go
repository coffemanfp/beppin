package utils

import (
	"database/sql"
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/lib/pq"
)

// UpdateProduct - Updates a product.
func UpdateProduct(db *sql.DB, productToUpdate, product models.Product) (err error) {
	identifier := productToUpdate.GetIdentifier()
	if identifier == nil {
		err = fmt.Errorf("failed to update product: %w (product)", errs.ErrNotProvidedOrInvalidObject)
		return
	}

	productToUpdate, err = SelectProduct(
		db,
		models.Product{
			ID: productToUpdate.ID,
		},
	)
	if err != nil {
		return
	}

	product = fillProductEmptyFields(product, productToUpdate)

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
		err = fmt.Errorf("failed to prepare the update (%v) product statement: %v", identifier, err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.Name,
		product.Description,
		pq.Array(product.Categories),
		productToUpdate.ID,
	)
	if err != nil {
		err = fmt.Errorf("failed to execute the update (%v) product statement: %v", identifier, err)
	}
	return
}

func fillProductEmptyFields(product, previousProductData models.Product) models.Product {

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
