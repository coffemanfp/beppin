package database

import "github.com/coffemanfp/beppin/database/models"

// ProductStorage reprensents all implementations for product utils.
type ProductStorage interface {
	CreateProduct(product models.Product) (int, error)
	GetProduct(productToFind models.Product) (models.Product, error)
	GetProducts(limit, offset int) (models.Products, error)
	UpdateProduct(productToUpdate, product models.Product) (int, error)
	DeleteProduct(product models.Product) (int, error)
}
