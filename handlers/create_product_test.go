package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/labstack/echo"
)

var createProduct string = `
	{
		"userID": 1,
		"name": "Product name",
		"description": "Product description",
		"categories": [
			"food"
		]
	}
`

func TestCreateProduct(t *testing.T) {
	// Setup server
	e := echo.New()
	e.Logger.Debug()

	setJWTMiddleware(t, e)
	setStorage(t)

	e.POST("/", handlers.CreateProduct)

	// Now the request
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createProduct))

	setAuthorizationRequest(t, req, token)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assertCreated(t, rec)
}
