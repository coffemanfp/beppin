package database

import (
	"fmt"

	"github.com/coffemanfp/beppin/database/models"
	dbu "github.com/coffemanfp/beppin/database/utils"
	errs "github.com/coffemanfp/beppin/errors"
)

func (dS defaultStorage) CreateProduct(product models.Product) (id int, err error) {
	exists, err := dS.ExistsUser(models.User{ID: product.UserID})
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) user: %w", product.UserID, errs.ErrNotExistentObject)
		return
	}

	id, err = dbu.InsertProduct(dS.db, product)
	return
}

func (dS defaultStorage) GetProduct(productToFind models.Product) (product models.Product, err error) {
	product, err = dbu.SelectProduct(dS.db, productToFind)
	return
}

func (dS defaultStorage) GetProducts(limit, offset int) (products models.Products, err error) {
	products, err = dbu.SelectProducts(dS.db, limit, offset)
	return
}

func (dS defaultStorage) UpdateProduct(productToUpdate, product models.Product) (id int, err error) {
	productToUpdate, err = dS.GetProduct(
		models.Product{
			ID: productToUpdate.ID,
		},
	)
	if err != nil {
		return
	}

	product = fillProductEmptyFields(product, productToUpdate)

	id, err = dbu.UpdateProduct(dS.db, productToUpdate, product)
	return
}

func (dS defaultStorage) DeleteProduct(productToDelete models.Product) (id int, err error) {
	id, err = dbu.DeleteProduct(dS.db, productToDelete)
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

	if product.Price == 0 {
		product.Price = previousProductData.Price
	}

	return product
}
