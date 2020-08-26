package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Setup server
	e := echo.New()

	e.Logger.Debug()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.Nil(t, handlers.GetProducts(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	var m models.ResponseMessage

	assert.Nil(t, json.NewDecoder(rec.Body).Decode(&m))
	assert.NotEmpty(t, m.ContentType)
	assert.Equal(t, models.TypeProducts, m.ContentType)
}

func TestFailedGetProducts(t *testing.T) {
	t.Parallel()
	limitTests := []struct {
		Name  string
		Limit string
	}{
		{
			Name:  "limit_negative_number",
			Limit: "-1",
		},
		{
			Name:  "limit_super_negative_number",
			Limit: "-986544567890",
		},
		{
			Name:  "limit_letters",
			Limit: "a",
		},
		{
			Name:  "limit_super_letters",
			Limit: "ajhkklaskldjkasksjdlfkjsdlfkjlasdkjfljsdf",
		},
		{
			Name:  "limit_super_greater_max",
			Limit: strconv.Itoa(int(config.GetSettings().MaxElementsPerPagination)) + "09876545678909876545678987678",
		},
	}

	offsetTests := []struct {
		Name   string
		Offset string
	}{
		{
			Name:   "offset_negative_number",
			Offset: "-1",
		},
		{
			Name:   "offset_super_negative_number",
			Offset: "-986544567890",
		},
		{
			Name:   "offset_letters",
			Offset: "a",
		},
		{
			Name:   "offset_super_letters",
			Offset: "ajhkklaskldjkasksjdlfkjsdlfkjlasdkjfljsdf",
		},
	}

	for _, ts := range limitTests {
		ts := ts
		t.Run(ts.Name, func(t *testing.T) {
			t.Parallel()
			// Setup server
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			query := req.URL.Query()
			query.Add("limit", ts.Limit)

			req.URL.RawQuery = query.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handlers.GetProducts(c)

			assertInvalidParam(t, "limit", err)
		})
	}

	for _, ts := range offsetTests {
		ts := ts
		t.Run(ts.Name, func(t *testing.T) {
			t.Parallel()
			// Setup server
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			query := req.URL.Query()

			query.Add("offset", ts.Offset)

			req.URL.RawQuery = query.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handlers.GetProducts(c)

			assertInvalidParam(t, "offset", err)
		})
	}
}
