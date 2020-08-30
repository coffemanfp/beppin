package handlers

import (
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// GetUsers - Get user.
func GetUsers(c echo.Context) (err error) {
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")

	var m models.ResponseMessage

	limit, err := utils.Atoi(limitParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: limit", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the limit param is exceeded, is setted to the default limit.
	m.LimitParamExceeded(&limit)

	// If the limit is not provided, is setted to the default limit.
	m.NotLimitParamProvided(&limit)

	offset, err := utils.Atoi(offsetParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: offset", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbUsers, err := Storage.GetUsers(limit, offset)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	var users models.Users

	if dbUsers == nil {
		users = make(models.Users, 0)
	} else {
		users = helpers.ShouldParseDBModelToModel(dbUsers).(models.Users)
	}

	m.Content = users
	m.ContentType = models.TypeUsers

	if m.Message == "" {
		m.Message = "Ok."
	}

	return c.JSON(http.StatusOK, m)
}
