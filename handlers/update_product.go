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

// UpdateProduct - Updates a product.
func UpdateProduct(c echo.Context) (err error) {
	productIDParam := c.Param("id")
	var m models.ResponseMessage

	userIDToken := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User.ID

	productID, err := utils.Atoi(productIDParam)
	if err != nil || productID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var product models.Product

	if err = c.Bind(&product); err != nil {
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

	updatedProduct, err := Storage.UpdateProduct(
		models.Product{
			ID: int64(productID),
		},
		product,
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: product", errs.ErrNotExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Updated."
	m.Content = updatedProduct
	m.ContentType = models.TypeProduct
	return c.JSON(http.StatusOK, m)
}
