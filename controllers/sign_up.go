package controllers

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
)

// SignUp - Register a user.
func SignUp(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = "invalid body"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !user.Validate() {
		m.Error = "invalid body"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbUserI, err := helpers.ParseModelToDBModel(user)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUser := dbUserI.(dbm.User)

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.InsertUser(db, dbUser)
	if err != nil {
		if err.Error() == errors.ErrExistentObject {
			m.Error = err.Error() + " (user)"

			return echo.NewHTTPError(http.StatusConflict, m)

		} else if err.Error() == errors.ErrNotExistentObject {
			m.Error = err.Error() + " (language)"

			return echo.NewHTTPError(http.StatusNotFound, m)
		}

		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Created."

	return c.JSON(http.StatusCreated, m)
}
