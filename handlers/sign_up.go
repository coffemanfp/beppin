package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
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

	if !user.Validate("signup") {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	newUser, err := Storage.SignUp(
		user,
	)
	if err != nil {
		unwrappedErr := errors.Unwrap(err)

		switch unwrappedErr {
		case errs.ErrExistentObject:
			m.Error = fmt.Sprintf("%v: user", errs.ErrExistentObject)
			return echo.NewHTTPError(http.StatusConflict, m)

		case errs.ErrNotExistentObject:
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)
			return echo.NewHTTPError(http.StatusNotFound, m)

		default:
			c.Logger().Error(err)
			m.Error = http.StatusText(http.StatusInternalServerError)

			return echo.NewHTTPError(http.StatusInternalServerError, m)
		}

	}

	claim := models.Claim{
		User: newUser,
	}

	token, err := claim.GenerateJWT()
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Created."
	m.Content = echo.Map{
		"token": token,
	}
	m.ContentType = models.TypeToken
	return c.JSON(http.StatusCreated, m)
}
