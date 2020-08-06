package controllers

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
	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// UpdateUser - Updates a user.
func UpdateUser(c echo.Context) (err error) {
	userIDParam := c.Param("id")
	var m models.ResponseMessage

	userID, err := utils.Atoi(userIDParam)
	if err != nil || userID == 0 {
		m.Error = fmt.Sprintf("%v: %s", errs.ErrInvalidParam, "id")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbuserI, err := helpers.ParseModelToDBModel(user)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUser := dbuserI.(dbm.User)

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.UpdateUser(db, userID, "", dbUser)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: %s", errs.ErrExistentObject, "user")
			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Updated."

	return c.JSON(http.StatusOK, m)
}
