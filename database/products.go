package database

import (
	"fmt"

	dbu "github.com/coffemanfp/beppin/database/utils"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

func (dS defaultStorage) CreateProduct(product models.Product) (createdProduct models.Product, err error) {
	tx, err := dS.db.Begin()
	if err != nil {
		return
	}

	exists, err := dbu.ExistsUser(tx, models.User{ID: product.UserID})
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) user: %w", product.UserID, errs.ErrNotExistentObject)
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	createdProduct, err = dbu.InsertProduct(tx, product)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	for _, file := range product.Images {
		exists, err = dbu.ExistsFile(tx, models.File{ID: file.ID})
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		if !exists {
			err = fmt.Errorf("failed to check (%d) file: %w", file.ID, errs.ErrNotExistentObject)
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		err = dbu.InsertProductFile(tx, createdProduct.ID, file.ID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}
	}
	for _, category := range product.Categories {
		exists, err = dbu.ExistsCategory(tx, models.Category{ID: category.ID})
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		if !exists {
			err = fmt.Errorf("failed to check (%d) category: %w", category.ID, errs.ErrNotExistentObject)
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		err = dbu.InsertProductCategory(tx, createdProduct.ID, category.ID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
	}
	return
}

func (dS defaultStorage) AddProductCategory(productID int64, categoryID int64) (err error) {
	exists, err := dbu.ExistsProduct(dS.db, models.Product{ID: productID})
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) : %w", productID, errs.ErrNotExistentObject)
		return
	}

	exists, err = dbu.ExistsCategory(dS.db, models.Category{ID: categoryID})
	if err != nil {
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) category: %w", categoryID, errs.ErrNotExistentObject)
		return
	}

	exists, err = dbu.ExistsProductCategory(dS.db, productID, categoryID)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to check product_category: %w", errs.ErrExistentObject)
		return
	}

	err = dbu.InsertProductCategory(dS.db, productID, categoryID)
	return
}

func (dS defaultStorage) AddProductCategories(productID int64, categories models.Categories) (err error) {
	tx, err := dS.db.Begin()
	if err != nil {
		return
	}

	exists, err := dbu.ExistsProduct(tx, models.Product{ID: productID})
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) : %w", productID, errs.ErrNotExistentObject)
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	for _, category := range categories {
		exists, err = dbu.ExistsCategory(tx, category)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		if !exists {
			err = fmt.Errorf("failed to check (%d) category: %w", category.ID, errs.ErrNotExistentObject)
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		exists, err = dbu.ExistsProductCategory(tx, productID, category.ID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		if exists {
			err = fmt.Errorf("failed to check product_category: %w", errs.ErrExistentObject)
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		err = dbu.InsertProductCategory(tx, productID, category.ID)
		if err != nil {
			fmt.Println("error aqui")
			err2 := tx.Rollback()
			if err2 != nil {
				fmt.Println("doble error")
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
	}
	return
}

func (dS defaultStorage) GetProduct(productToFind models.Product) (product models.Product, err error) {
	product, err = dbu.SelectProduct(dS.db, productToFind)
	if err != nil {
		return
	}

	files, err := dbu.SelectProductFiles(dS.db, productToFind)
	if err != nil {
		return
	}

	categories, err := dbu.SelectProductCategories(dS.db, productToFind)
	if err != nil {
		return
	}

	product.Images = files
	product.Categories = categories
	return
}

func (dS defaultStorage) GetProducts(limit, offset int) (products models.Products, err error) {
	products, err = dbu.SelectProducts(dS.db, limit, offset)

	var files models.Files
	var categories models.Categories
	for i := 0; i < len(products); i++ {
		files, err = dbu.SelectProductFiles(dS.db, products[i])
		if err != nil {
			return
		}

		categories, err = dbu.SelectProductCategories(dS.db, products[i])
		if err != nil {
			return
		}

		products[i].Images = files
		products[i].Categories = categories
	}
	return
}

func (dS defaultStorage) UpdateProduct(productToUpdate, product models.Product) (updatedProduct models.Product, err error) {
	updatedProduct, err = dbu.UpdateProduct(dS.db, productToUpdate, product)
	return
}

func (dS defaultStorage) UpdateProductCategories(productID int64, categories models.Categories) (err error) {
	tx, err := dS.db.Begin()
	if err != nil {
		return
	}

	exists, err := dbu.ExistsProduct(tx, models.Product{ID: productID})
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	if !exists {
		err = fmt.Errorf("failed to check (%d) : %w", productID, errs.ErrNotExistentObject)
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	for _, category := range categories {
		exists, err = dbu.ExistsCategory(tx, category)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}

		if !exists {
			err = fmt.Errorf("failed to check (%d) category: %w", category.ID, errs.ErrNotExistentObject)
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}
	}

	err = dbu.DeleteProductCategories(tx, productID)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
		return
	}

	for _, category := range categories {
		err = dbu.InsertProductCategory(tx, productID, category.ID)
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				err = fmt.Errorf("%s\n%s", err2, err)
			}
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			err = fmt.Errorf("%s\n%s", err2, err)
		}
	}
	return
}

func (dS defaultStorage) DeleteProductCategories(productID int64) (err error) {
	err = dbu.DeleteProductCategories(dS.db, productID)
	return
}

func (dS defaultStorage) DeleteProductCategory(productID int64, categoryID int64) (err error) {
	count, err := dbu.SelectProductCategoriesCount(dS.db, productID)
	if err != nil {
		return
	}

	if count <= 1 {
		err = fmt.Errorf("failed to delete product: %w. the product does not have enough categories", errs.ErrInvalidData)
		return
	}

	err = dbu.DeleteProductCategory(dS.db, productID, categoryID)
	return
}

func (dS defaultStorage) DeleteProduct(productToDelete models.Product) (id int, err error) {
	id, err = dbu.DeleteProduct(dS.db, productToDelete)
	return
}
