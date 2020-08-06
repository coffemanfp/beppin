package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// DeleteProduct - Delete a product.
func DeleteProduct(c echo.Context) (err error) {
	var m models.ResponseMessage
	var productID int

	productIDParam := c.Param("id")

	if productID, err = utils.Atoi(productIDParam); err != nil || productID == 0 {
		m.Error = fmt.Sprintf("%v: %s", errs.ErrInvalidParam, "id")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.DeleteProduct(db, productID)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: %s", errs.ErrNotExistentObject, "product")

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Deleted."
	return c.JSON(http.StatusOK, m)
}
