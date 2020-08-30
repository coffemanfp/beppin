package handlers

import (
	"fmt"
	"net/http"

	dbm "github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
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
		dbm.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		},
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
		User: helpers.ShouldParseDBModelToModel(dbUser).(models.User),
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
