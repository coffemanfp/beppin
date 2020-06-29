package controllers

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// GetProducts - Get products.
func GetProducts(c echo.Context) (err error) {
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	var m models.ResponseMessage

	var limit, offset int

	limit, err = utils.Atoi(limitParam)
	if err != nil {
		m.Error = "limit param not valid"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	offset, err = utils.Atoi(offsetParam)
	if err != nil {
		m.Error = "offset param not valid"

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
