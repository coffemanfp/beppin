package handlers

import (
	"errors"
	"fmt"
	"net/http"

	dbm "github.com/coffemanfp/beppin-server/database/models"
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

	dbProduct, err := Storage.GetProduct(
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
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Content = helpers.ShouldParseDBModelToModel(dbProduct).(models.Product)
	return c.JSON(http.StatusOK, m)
}
