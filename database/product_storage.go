package database

import "github.com/coffemanfp/beppin/models"

// ProductStorage reprensents all implementations for product utils.
type ProductStorage interface {
	CreateProduct(product models.Product) (models.Product, error)
	AddProductCategory(productID int64, categoryID int64) (err error)
	AddProductCategories(productID int64, categories models.Categories) (err error)
	GetProduct(productToFind models.Product) (models.Product, error)
	GetProducts(limit, offset int) (models.Products, error)
	UpdateProduct(productToUpdate, product models.Product) (models.Product, error)
	UpdateProductCategories(productID int64, categories models.Categories) error
	DeleteProductCategories(productID int64) error
	DeleteProductCategory(productID int64, categoryID int64) error
	DeleteProduct(product models.Product) (int, error)
}
