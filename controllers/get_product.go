package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// GetProduct - Get a product.
func GetProduct(c echo.Context) (err error) {
	var m models.ResponseMessage
	var productID int

	productIDParam := c.Param("id")

	if productID, err = utils.Atoi(productIDParam); err != nil || productID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbProduct, err := dbu.SelectProduct(
		db,
		dbm.Product{
			ID: int64(productID),
		},
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: product", errs.ErrExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	productI, err := helpers.ParseDBModelToModel(dbProduct)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	product := productI.(models.Product)

	m.Content = product
	return c.JSON(http.StatusOK, m)
}
