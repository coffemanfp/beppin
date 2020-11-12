package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/coffemanfp/beppin/config"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UpdateAvatar Updates the user avatar.
func UpdateAvatar(c echo.Context) (err error) {
	var m models.ResponseMessage
	var userID int

	user := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User
	userIDParam := c.Param("id")

	if userID, err = utils.Atoi(userIDParam); err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the user to update isn't equal to the current user,
	// is not authorized
	if userID != int(user.ID) {
		m.Error = http.StatusText(http.StatusUnauthorized)

		return echo.NewHTTPError(http.StatusUnauthorized, m)
	}

	// Getting image from form input
	file, err := c.FormFile("file")
	if err != nil {
		m.Error = fmt.Sprintf("%v: file", errs.ErrInvalidData)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// If the avatar image is greater than 800kb, it's not valid
	if file.Size > 800000 {
		m.Error = fmt.Sprintf("%v: file", errs.ErrInvalidData)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	src, err := file.Open()
	if err != nil {
		m.Error = fmt.Sprintf("%v: file", errs.ErrInvalidData)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}
	defer src.Close()

	// Destination
	var dest *os.File

	destPath := path.Join(
		config.GlobalSettings.Assets,
		"avatars",
		strconv.Itoa(userID)+path.Ext(file.Filename))

	if exists, _ := utils.ExistsFile(destPath); exists {
		dest, err = os.OpenFile(destPath, os.O_WRONLY, 0777)
	} else {
		dest, err = os.Create(destPath)
	}
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	dbUser, err := Storage.UpdateUser(
		models.User{
			ID: int64(userID),
		},
		models.User{
			Avatar: &models.Avatar{
				URL: destPath,
			},
		},
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
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

	m.Message = "Updated."
	m.Content = echo.Map{
		"token": token,
	}
	m.ContentType = models.TypeToken
	return c.JSON(http.StatusOK, m)
}
