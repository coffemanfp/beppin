package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UpdateUser - Updates a user.
func UpdateUser(c echo.Context) (err error) {
	var m models.ResponseMessage

	userIDParam := c.Param("id")
	userIDToken := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User.ID

	userID, err := utils.Atoi(userIDParam)
	if err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the user to update isn't equal to the current user,
	// is not authorized
	if userID != int(userIDToken) {
		m.Error = http.StatusText(http.StatusUnauthorized)

		return echo.NewHTTPError(http.StatusUnauthorized, m)
	}

	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbUser, err := Storage.UpdateUser(
		models.User{
			ID: int64(userID),
		},
		user,
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)
			return echo.NewHTTPError(http.StatusNotFound, m)
		} else if errors.Is(err, errs.ErrExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrExistentObject)
			return echo.NewHTTPError(http.StatusConflict, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	claim := models.Claim{
		User: models.User{
			ID:       dbUser.ID,
			Language: dbUser.Language,
			Avatar:   dbUser.Avatar,
			Username: dbUser.Username,
			Theme:    dbUser.Theme,
			Currency: dbUser.Currency,
		},
	}

	token, err := claim.GenerateJWT()
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Updated."
	m.Content = echo.Map{
		"token": token,
	}
	m.ContentType = models.TypeToken
	return c.JSON(http.StatusOK, m)
}
