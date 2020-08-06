package controllers

import (
	"fmt"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
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

	var limit, offset int

	limit, err = utils.Atoi(limitParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: %s", errs.ErrInvalidParam, "limit")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the limit param is exceeded, is setted to the default limit.
	m.LimitParamExceeded(&limit)

	// If the limit is not provided, is setted to the default limit.
	m.NotLimitParamProvided(&limit)

	offset, err = utils.Atoi(offsetParam)
	if err != nil {
		m.Error = fmt.Sprintf("%v: %s", errs.ErrInvalidParam, "offset")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUsers, err := dbu.SelectUsers(db, limit, offset)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	var users models.Users

	if dbUsers == nil {
		users = make(models.Users, 0)
	} else {
		usersI, err := helpers.ParseDBModelToModel(dbUsers)
		if err != nil {
			c.Logger().Error(err)

			return echo.ErrInternalServerError
		}

		users = usersI.(models.Users)
	}

	m.Content = users

	if m.Message == "" {
		m.Message = "Ok."
	}

	return c.JSON(http.StatusOK, m)
}
