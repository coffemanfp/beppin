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

// SignUp - Register a user.
func SignUp(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !user.Validate() {
		m.Error = errs.ErrInvalidBody

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
		unwrappedErr := errors.Unwrap(err)

		switch unwrappedErr {
		case errs.ErrExistentObject:
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)
			return echo.NewHTTPError(http.StatusConflict, m)

		case errs.ErrNotExistentObject:
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)
			return echo.NewHTTPError(http.StatusNotFound, m)

		default:
			c.Logger().Error(err)

			return echo.ErrInternalServerError
		}

	}

	m.Message = "Created."
	return c.JSON(http.StatusCreated, m)
}
