package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
)

// CreateUser - Creates a user.
func CreateUser(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = "invalid body"
		fmt.Println(time.Now().String())

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !user.Validate() {
		m.Error = "invalid body"
		fmt.Println("aqui 2")

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	dbUserI, err := helpers.ParseModelToDBModel(user)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUser := dbUserI.(dbm.User)

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.InsertUser(db, dbUser)
	if err != nil {
		if err.Error() == errors.ErrExistentObject {
			m.Error = err.Error() + " (user)"

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Created."

	return c.JSON(http.StatusCreated, m)
}
