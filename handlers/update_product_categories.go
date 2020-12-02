package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UpdateProductCategories - Updates all the categories related with a product.
func UpdateProductCategories(c echo.Context) (err error) {
	productIDParam := c.Param("id")
	var m models.ResponseMessage

	userIDToken := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User.ID

	productID, err := utils.Atoi(productIDParam)
	if err != nil || productID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var categories models.Categories

	if err = c.Bind(&categories); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Check if it even has a value
	if len(categories) < 1 {
		m.Message = "A product must have at least a category."
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Omit any value repeated
	categoriesMap := make(map[int64]bool)

	for _, category := range categories {
		categoriesMap[category.ID] = true
	}

	categories = models.Categories{}
	for categoryID := range categoriesMap {
		categories = append(categories, models.Category{ID: categoryID})
	}

	// Get old product info
	oldProduct, err := Storage.GetProduct(models.Product{ID: int64(productID)})
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: product", errs.ErrNotExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	if oldProduct.UserID != userIDToken {
		m.Error = http.StatusText(http.StatusUnauthorized)

		return echo.NewHTTPError(http.StatusUnauthorized, m)
	}

	err = Storage.UpdateProductCategories(
		int64(productID),
		categories,
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = err.Error()

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Updated."
	m.Content = models.Product{
		ID:         int64(productID),
		Categories: categories,
	}
	m.ContentType = models.TypeProduct
	return c.JSON(http.StatusOK, m)
}
