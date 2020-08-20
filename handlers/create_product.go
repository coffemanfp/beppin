package handlers

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
	"github.com/labstack/echo"
)

// CreateProduct - Creates a product.
func CreateProduct(c echo.Context) (err error) {
	var m models.ResponseMessage
	var product models.Product

	if err = c.Bind(&product); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !product.Validate() {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbProductI, err := helpers.ParseModelToDBModel(product)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbProduct := dbProductI.(dbm.Product)

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.InsertProduct(db, dbProduct)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Created."

	return c.JSON(http.StatusCreated, m)
}
