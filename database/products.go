package database

import (
	"fmt"

	dbu "github.com/coffemanfp/beppin/database/utils"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
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
	id, err = dbu.UpdateProduct(dS.db, productToUpdate, product)
	return
}

func (dS defaultStorage) DeleteProduct(productToDelete models.Product) (id int, err error) {
	id, err = dbu.DeleteProduct(dS.db, productToDelete)
	return
}
