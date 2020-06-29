package controllers

import (
	"errors"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// UpdateProduct - Updates a product.
func UpdateProduct(c echo.Context) (err error) {
	productIDParam := c.Param("id")
	var m models.ResponseMessage

	productID, err := utils.Atoi(productIDParam)
	if err != nil || productID == 0 {
		m.Error = errors.New("id param not valid")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var product models.Product

	if err = c.Bind(&product); err != nil {
		m.Error = errors.New("invalid body")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !product.ValidateUpdate() {
		m.Error = errors.New("invalid product fields")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.UpdateProduct(db, productID, product)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Updated."

	return c.JSON(http.StatusOK, m)
}