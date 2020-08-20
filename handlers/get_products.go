package handlers

import (
	"fmt"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
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
		m.Error = fmt.Sprintf("%v: limit", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the limit param is exceeded, is setted to the default limit.
	m.LimitParamExceeded(&limit)

	// If the limit is not provided, is setted to the default limit.
	m.NotLimitParamProvided(&limit)

	offset, err = utils.Atoi(offsetParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: offset", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbProducts, err := db.GetProducts(limit, offset)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	var products models.Products

	if dbProducts == nil {
		products = make(models.Products, 0)
	} else {
		productsI, err := helpers.ParseDBModelToModel(dbProducts)
		if err != nil {
			c.Logger().Error(err)

			return echo.ErrInternalServerError
		}

		products = productsI.(models.Products)
	}

	m.Content = products

	if m.Message == "" {
		m.Message = "Ok."
	}

	return c.JSON(http.StatusOK, m)
}
