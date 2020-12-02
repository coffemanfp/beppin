package database

import "github.com/coffemanfp/beppin/models"

// CategoriesStorage reprensents all implementations for categories utils.
type CategoriesStorage interface {
	GetCategories() (models.Categories, error)
	GetCategory(categoryToFind models.Category) (models.Category, error)
}
