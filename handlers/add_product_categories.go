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

// AddProductCategories - Adds a product category relation.
func AddProductCategories(c echo.Context) (err error) {
	var m models.ResponseMessage

	userIDToken := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User.ID
	productIDParam := c.Param("id")

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

	// Omit any value repeated
	categoriesMap := make(map[int64]bool)

	for _, category := range categories {
		categoriesMap[category.ID] = true
	}

	categories = models.Categories{}
	for categoryID := range categoriesMap {
		categories = append(categories, models.Category{ID: categoryID})
	}

	err = Storage.AddProductCategories(int64(productID), categories)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) ||
			errors.Is(err, errs.ErrExistentObject) {
			m.Error = err.Error()

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Created."
	return c.JSON(http.StatusCreated, m)
}
