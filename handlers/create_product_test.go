package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/models"
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

func TestFailedCreateProduct(t *testing.T) {
	tests := []struct {
		Name string
		Body interface{}
	}{
		{
			Name: "invalid_body",
			Body: "alksdlkadjs",
		},
		{
			Name: "empty_product",
			Body: models.Product{},
		},
		{
			Name: "not_existent_user",
			Body: models.Product{
				UserID:      time.Now().Unix(),
				Name:        exampleProducts[0].Name,
				Description: exampleProducts[0].Description,
				Categories:  exampleProducts[0].Categories,
			},
		},
	}

	for _, ts := range tests {
		t.Run(ts.Name, func(t *testing.T) {
			// Setup server
			e := echo.New()
			e.Logger.Debug()

			setJWTMiddleware(t, e)

			e.POST("/", handlers.CreateProduct)

			bodyJSON, err := json.Marshal(ts.Body)
			assert.Nil(t, err)

			// Now the request
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyJSON))

			setAuthorizationRequest(t, req, token)

			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			if ts.Name == "not_existent_user" {
				assertNotExistent(t, rec, "user")
			} else {
				assertInvalidBody(t, rec)
			}
		})
	}
}
