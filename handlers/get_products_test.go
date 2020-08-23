package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/beppin-server/config"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Setup server
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.Nil(t, handlers.GetProducts(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	var m models.ResponseMessage
	assert.Nil(t, json.NewDecoder(rec.Body).Decode(&m))

	products, ok := m.Content.([]interface{})
	if !ok {
		t.Fatalf(
			"response message content type not expected: expected (%T), gotted (%T)",
			[]interface{}{},
			m.Content,
		)
	}

	for _, product := range products {
		if _, ok = m.Content.([]interface{}); !ok {
			if !ok {
				t.Fatalf(
					"response message content type not expected: expected (%T), gotted (%T)",
					models.Product{},
					product,
				)
			}
		}
	}

	if len(products) > int(config.GetSettings().MaxElementsPerPagination) {
		t.Fatalf(
			"max elements per pagination exceded: max expected (%d), gotted (%d)",
			config.GetSettings().MaxElementsPerPagination,
			len(products),
		)
	}
}

func TestFailedGetProducts(t *testing.T) {
	// Setup server
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	query := req.URL.Query()

	t.Run("limit_validation", func(t *testing.T) {
		query.Add("limit", "-1")

		req.URL.RawQuery = query.Encode()

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handlers.GetProducts(c)
		assert.NotNil(t, err)

		assert.Contains(t, err.Error(), errs.ErrInvalidParam)
		assert.Contains(t, err.Error(), "limit")
	})
}
