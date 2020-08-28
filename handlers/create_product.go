package handlers

import (
	"errors"
	"fmt"
	"net/http"

	dbm "github.com/coffemanfp/beppin-server/database/models"
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
		fmt.Println(err)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !product.Validate() {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	err = Storage.CreateProduct(
		helpers.ShouldParseModelToDBModel(product).(dbm.Product),
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Created."
	return c.JSON(http.StatusCreated, m)
}
