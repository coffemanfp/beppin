package controllers

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// GetUser - Get a user.
func GetUser(c echo.Context) (err error) {
	var m models.ResponseMessage
	var userID int

	userIDParam := c.Param("id")

	if userID, err = utils.Atoi(userIDParam); err != nil || userID == 0 {
		m.Error = "id param not valid"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbuser, err := dbu.SelectUser(db, userID, "")
	if err != nil {
		if err.Error() == errors.ErrNotExistentObject {
			m.Error = err.Error() + " (user)"
			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	userI, err := helpers.ParseDBModelToModel(dbuser)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	user := userI.(models.User)

	m.Content = user

	return c.JSON(http.StatusOK, m)
}
