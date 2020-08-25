package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

var token string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJsYW5ndWFnZSI6ImVzLUVTIiwidXNlcm5hbWUiOiJjb2ZmZW1hbmZwIiwidGhlbWUiOiJsaWdodCJ9fQ.GJcykxeN4yfE7CVi1xu4zVYstPgODCuNtrgq4T11gA4"

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

func assertCreated(t *testing.T, rec *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusCreated, rec.Code)

	var m models.ResponseMessage
	assert.Nil(t, json.NewDecoder(rec.Body).Decode(&m))
	assert.Equal(t, "Created.", m.Message)
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
