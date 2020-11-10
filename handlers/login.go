package handlers

import (
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/labstack/echo"
)

// Login - Login user.
func Login(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !user.Validate("login") {
		m.Error = fmt.Sprintf("%v", errs.ErrInvalidUserLogin)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbUser, match, err := Storage.Login(
		user,
	)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	if !match {
		return echo.ErrUnauthorized
	}

	claim := models.Claim{
		User: dbUser,
	}

	token, err := claim.GenerateJWT()
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Ok."
	m.Content = echo.Map{
		"token": token,
	}
	m.ContentType = models.TypeToken

	return c.JSON(http.StatusOK, m)
}
