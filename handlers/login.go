package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login - Login user.
func Login(c echo.Context) (err error) {
	var m models.ResponseMessage
	var user models.User

	if err = c.Bind(&user); err != nil {
		m.Error = errs.ErrInvalidBody
		fmt.Println(err)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	if !user.ValidateLogin() {
		m.Error = fmt.Sprintf("%v", errs.ErrInvalidUserLogin)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	dbUser, match, err := db.Login(
		dbm.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		},
	)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	if !match {
		return echo.ErrUnauthorized
	}

	userI, err := helpers.ParseDBModelToModel(dbUser)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	user = userI.(models.User)

	fmt.Println(user)

	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token, err := claim.GenerateJWT()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Ok."
	m.Content = echo.Map{
		"token": token,
	}

	return c.JSON(http.StatusOK, m)
}
