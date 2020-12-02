package database

import (
	dbu "github.com/coffemanfp/beppin/database/utils"
	"github.com/coffemanfp/beppin/models"
)

func (dS defaultStorage) GetCategories() (categories models.Categories, err error) {
	categories, err = dbu.SelectCategories(dS.db)
	return
}

func (dS defaultStorage) GetCategory(categoryToFind models.Category) (category models.Category, err error) {
	category, err = dbu.SelectCategory(dS.db, categoryToFind)
	return
}
