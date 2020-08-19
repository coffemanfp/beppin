package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// UpdateAvatar Updates the user avatar.
func UpdateAvatar(c echo.Context) (err error) {
	var avatar models.Avatar
	var m models.ResponseMessage
	var userID int

	userIDParam := c.Param("id")

	if userID, err = utils.Atoi(userIDParam); err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if err = c.Bind(&avatar); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !avatar.Validate() {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	avatarURL := avatar.URL

	var db *sql.DB
	db, err = database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	exists, err := dbu.ExistsUser(db, dbm.User{ID: int64(userID)})
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	if !exists {
		m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)

		return echo.NewHTTPError(http.StatusNotFound, m)
	}

	// If the user will save their avatar in our file system
	if avatar.Data != "" && avatar.URL == "" {
		avatarURL, err = avatar.Save(strconv.Itoa(userID))
		if err != nil {
			c.Logger().Error(err)

			return echo.ErrInternalServerError
		}
	}

	err = dbu.UpdateAvatar(
		db,
		avatarURL,
		dbm.User{ID: int64(userID)},
	)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	return
}
