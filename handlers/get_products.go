package handlers

import (
	"fmt"
	"net/http"

	"github.com/coffemanfp/beppin-server/config"
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
	maxElementsPerPagination := config.GlobalSettings.MaxElementsPerPagination

	var m models.ResponseMessage

	limit, err := utils.Atoi(limitParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: limit", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if limit < 0 {
		m.Error = fmt.Sprintf("%v: limit", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the limit is not provided, is setted to the default limit.
	if limit == 0 {
		limit = maxElementsPerPagination
		m.Message = fmt.Sprintf("%s: setted to %d",
			models.MessageNotLimitParam,
			maxElementsPerPagination,
		)

		// If the limit param is exceeded, is setted to the default limit.
	} else if limit > maxElementsPerPagination {
		limit = maxElementsPerPagination
		m.Message = fmt.Sprintf("%s: setted to %d",
			models.MessageLimitParamExceeded,
			maxElementsPerPagination,
		)
	}

	offset, err := utils.Atoi(offsetParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: offset", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if offset < 0 {
		m.Error = fmt.Sprintf("%v: offset", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbProducts, err := Storage.GetProducts(limit, offset)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	var products models.Products

	if dbProducts == nil {
		products = make(models.Products, 0)
	} else {
		products = helpers.ShouldParseDBModelToModel(dbProducts).(models.Products)
	}

	m.Content = products
	m.ContentType = models.TypeProducts

	if m.Message == "" {
		m.Message = "Ok."
	}

	return c.JSON(http.StatusOK, m)
}
