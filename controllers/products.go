package controllers

import (
	"net/http"
	"strconv"

	"github.com/coffemanfp/beppin-server/database"
	"github.com/labstack/echo"
)

// GetProducts - Get products.
func GetProducts(c *echo.Context) (err error) {
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	var limit, offset int

	if limitParam == "" {
		limitParam = "0"
	}
	if offsetParam == "" {
		offsetParam = "0"
	}

	if limit, err = strconv.Atoi(limitParam); err != nil {
		return echo.ErrInternalServerError
	}

	if offset, err = strconv.Atoi(offsetParam); err != nil {
		return echo.ErrInternalServerError
	}

	db, err := database.GetConn()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	products, err := dbu.GetProducts(limit, offset)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, products)
}
