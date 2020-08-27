package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// Setup server
	e := echo.New()
	e.Logger.Debug()

	setJWTMiddleware(t, e)
	setStorage(t)

	if !existsLanguage(t, exampleLanguage) {
		insertLanguage(t, exampleLanguage)
	}
	if !existsUser(t, exampleUser) {
		insertUser(t, exampleUser)
	}

	e.POST("/", handlers.CreateProduct)

	productJSON, err := json.Marshal(exampleProducts[0])
	assert.Nil(t, err)

	// Now the request
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(productJSON))

	setAuthorizationRequest(t, req, token)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assertCreated(t, rec)
}
