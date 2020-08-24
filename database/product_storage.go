package database

import "github.com/coffemanfp/beppin-server/database/models"

// ProductStorage reprensents all implementations for product utils.
type ProductStorage interface {
	CreateProduct(product models.Product) error
	GetProduct(productToFind models.Product) (models.Product, error)
	GetProducts(limit, offset uint64) (models.Products, error)
	UpdateProduct(productToUpdate, product models.Product) error
	DeleteProduct(product models.Product) error
}
