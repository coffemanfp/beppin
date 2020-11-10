package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/labstack/echo"
)

// DeleteUser - Delete a user.
func DeleteUser(c echo.Context) (err error) {
	var m models.ResponseMessage
	var userID int

	userIDParam := c.Param("id")

	if userID, err = utils.Atoi(userIDParam); err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	id, err := Storage.DeleteUser(
		models.User{
			ID: int64(userID),
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

	m.Message = "Deleted."
	m.Content = models.User{
		ID: int64(id),
	}
	m.ContentType = models.TypeUser

	return c.JSON(http.StatusOK, m)
}
