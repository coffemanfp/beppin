package controllers

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// DeleteUser - Delete a user.
func DeleteUser(c echo.Context) (err error) {
	var m models.ResponseMessage
	var userID int

	userIDParam := c.Param("id")

	if userID, err = utils.Atoi(userIDParam); err != nil || userID == 0 {
		m.Error = "id param not valid"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.DeleteUser(db, userID, "")
	if err != nil {
		if err.Error() == errors.ErrNotExistentObject {
			m.Error = err.Error() + " (user)"
			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Deleted."

	return c.JSON(http.StatusOK, m)
}
