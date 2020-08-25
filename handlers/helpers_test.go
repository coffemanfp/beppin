package handlers_test

import (
	"net/http"
	"testing"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

func assertInvalidParam(t *testing.T, param string, err error) {
	t.Helper()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errs.ErrInvalidParam)
	assert.Contains(t, err.Error(), param)
}

func assertInternalServerError(t *testing.T, err error) {
	t.Helper()

	if assert.NotNil(t, err) {
		echoError, ok := err.(*echo.HTTPError)
		assert.Equal(t, true, ok)
		assert.Equal(t, http.StatusInternalServerError, echoError.Code)
		assert.Equal(t, echo.ErrInternalServerError.Message, echoError.Message)
	}
}

func setJWTMiddleware(t *testing.T, e *echo.Echo) {
	jwtConfig := middleware.JWTConfig{
		Claims:      &models.Claim{},
		SigningKey:  []byte(config.GetSettings().SecretKey),
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}

	e.Use(middleware.JWTWithConfig(jwtConfig))
}

func setAuthorizationRequest(t *testing.T, req *http.Request, token string) {
	t.Helper()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, token)
}
