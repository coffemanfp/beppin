package controllers

import (
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/stretchr/gomniauth"
)

// LoginWithProvider login the user with a provider.
func LoginWithProvider(c echo.Context) (err error) {
	var m models.ResponseMessage

	providerParam := c.Param("provider")
	if providerParam == "" {
		m.Error = fmt.Sprintf("%v: provider", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	provider, err := gomniauth.Provider(providerParam)
	if err != nil {
		err = fmt.Errorf("failed to get (%s) provider: %v", providerParam, err)
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		err = fmt.Errorf("failed to get begin auth url for (%s) provider: %v", providerParam, err)
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	c.Redirect(http.StatusTemporaryRedirect, loginURL)
	return
}
