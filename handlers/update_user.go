package handlers

import (
	"errors"
	"fmt"
	"net/http"

	dbm "github.com/coffemanfp/beppin/database/models"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/helpers"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/labstack/echo"
)

// UpdateUser - Updates a user.
func UpdateUser(c echo.Context) (err error) {
	userIDParam := c.Param("id")
	var m models.ResponseMessage

	userID, err := utils.Atoi(userIDParam)
	if err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	id, err := Storage.UpdateUser(
		dbm.User{
			ID: int64(userID),
		},
		helpers.ShouldParseModelToDBModel(user).(dbm.User),
	)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrExistentObject)
			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Updated."
	m.Content = models.User{
		ID: int64(id),
	}
	m.ContentType = models.TypeUser
	return c.JSON(http.StatusOK, m)
}
