package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	errs "github.com/coffemanfp/beppin-server/errors"
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

	assertResponseMessage(t, "Created.", decodeResponseMessage(t, rec))
}

func TestFailedCreateProduct(t *testing.T) {
	tests := []struct {
		Name               string
		Body               interface{}
		ExpectedStatusCode int
		ExpectedError      string
	}{
		{
			Name:               "invalid_body",
			Body:               "alksdlkadjs",
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      errs.ErrInvalidBody,
		},
		{
			Name:               "empty_product",
			Body:               models.Product{},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      errs.ErrInvalidBody,
		},
		{
			Name: "not_existent_user",
			Body: models.Product{
				UserID:      time.Now().Unix(),
				Name:        exampleProducts[0].Name,
				Description: exampleProducts[0].Description,
				Categories:  exampleProducts[0].Categories,
			},
			ExpectedStatusCode: http.StatusNotFound,
			ExpectedError:      fmt.Sprintf("%v: user", errs.ErrNotExistentObject),
		},
	}

	for _, ts := range tests {
		t.Run(ts.Name, func(t *testing.T) {
			// Setup server
			e := echo.New()
			e.Logger.Debug()

			setJWTMiddleware(t, e)

			e.POST("/", handlers.CreateProduct)

			// Now the request
			bodyJSON, err := json.Marshal(ts.Body)
			assert.Nil(t, err)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyJSON))

			setAuthorizationRequest(t, req, token)

			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assertResponseError(t, ts.ExpectedError, decodeResponseMessage(t, rec))
		})
	}
}
