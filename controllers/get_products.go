package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
)

// GetProducts - Get products.
func GetProducts(c echo.Context) (err error) {
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	m := models.ResponseMessage{}

	var limit, offset int

	if limitParam == "" {
		limitParam = "0"
	}
	if offsetParam == "" {
		offsetParam = "0"
	}

	if limit, err = strconv.Atoi(limitParam); err != nil {
		m.Error = errors.New("limit param not valid")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if offset, err = strconv.Atoi(offsetParam); err != nil {
		m.Error = errors.New("offset param not valid")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	products, err := dbu.SelectProducts(db, limit, offset)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	if products == nil {
		products = make([]models.Product, 0)
	}

	m.Content = products
	m.Message = "Ok."

	return c.JSON(http.StatusOK, m)
}
