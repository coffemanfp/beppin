package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/coffemanfp/beppin/config"
	"github.com/coffemanfp/beppin/database"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/handlers"
	"github.com/coffemanfp/beppin/models"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	tests := []struct {
		Name        string
		QueryParams url.Values
		WithData    bool
	}{
		{
			Name: "without_products",
		},
		{
			Name:     "with_products",
			WithData: true,
		},
		{
			Name: "with_limit_param",
			QueryParams: url.Values{
				"limit": []string{
					"10",
				},
			},
			WithData: true,
		},
	}

	for _, ts := range tests {

		t.Run(ts.Name, func(t *testing.T) {

			// Setup server
			e := echo.New()
			e.Logger.Debug()

			setStorage(t)

			e.GET("/", handlers.GetProducts)

			req := httptest.NewRequest(http.MethodGet, "/", nil)

			req.URL.RawQuery = ts.QueryParams.Encode()
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			var m models.ResponseMessage
			m = decodeResponseMessage(t, rec)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, models.TypeProducts, m.ContentType)

			var product models.Product

			decoder, err := mapstructure.NewDecoder(
				&mapstructure.DecoderConfig{
					Result:     &product,
					DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
				},
			)
			assert.Nil(t, err)

			err = decoder.Decode(m.Content.([]interface{})[0])
			assert.Nil(t, err)

			limitParam := ts.QueryParams.Get("limit")

			if limitParam != "" {
				limit, err := strconv.Atoi(limitParam)
				assert.Nil(t, err)

				assert.GreaterOrEqual(t, limit, len(m.Content.([]interface{})))
			}
		})
	}
}

func TestFailedGetProducts(t *testing.T) {
	t.Parallel()
	invalidParamLimit := fmt.Sprintf("%v: limit", errs.ErrInvalidParam)
	invalidParamOffset := fmt.Sprintf("%v: offset", errs.ErrInvalidParam)

	tests := []struct {
		Name               string
		QueryParams        url.Values
		ExpectedStatusCode int
		ExpectedError      string
		WithDatabase       bool
		WithProducts       bool
	}{
		{
			Name: "limit_negative_number",
			QueryParams: url.Values{
				"limit": []string{
					"-1",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamLimit,
		},
		{
			Name: "limit_super_negative_number",
			QueryParams: url.Values{
				"limit": []string{
					"-986544567890",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamLimit,
		},
		{
			Name: "limit_letters",
			QueryParams: url.Values{
				"limit": []string{
					"a",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamLimit,
		},
		{
			Name: "limit_super_letters",
			QueryParams: url.Values{
				"limit": []string{
					"ajhkklaskldjkasksjdlfkjsdlfkjlasdkjfljsdf",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamLimit,
		},
		{
			Name: "limit_super_greater_max",
			QueryParams: url.Values{
				"limit": []string{
					strconv.Itoa(config.GlobalSettings.MaxElementsPerPagination) + "09876545678909876545678987678",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamLimit,
		},
		{
			Name: "offset_negative_number",
			QueryParams: url.Values{
				"offset": []string{
					"-1",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamOffset,
		},
		{
			Name: "offset_super_negative_number",
			QueryParams: url.Values{
				"offset": []string{
					"-986544567890",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamOffset,
		},
		{
			Name: "offset_letters",
			QueryParams: url.Values{
				"offset": []string{
					"a",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamOffset,
		},
		{
			Name: "offset_super_letters",
			QueryParams: url.Values{
				"offset": []string{
					"ajhkklaskldjkasksjdlfkjsdlfkjlasdkjfljsdf",
				},
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedError:      invalidParamOffset,
		},
		{
			Name:               "without_database",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedError:      http.StatusText(http.StatusInternalServerError),
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run(ts.Name, func(t *testing.T) {
			t.Parallel()
			// Setup server
			e := echo.New()
			e.Logger.Debug()

			e.GET("/", handlers.GetProducts)

			if ts.WithDatabase {
				var storage database.Storage
				storage, err := database.NewDefault()

				assert.Nil(t, err)

				handlers.Storage = storage

				// if ts.WithProducts {
				// 	if !existsLanguage(t, exampleLanguage) {
				// 		insertLanguage(t, exampleLanguage)
				// 	}
				// 	if !existsUser(t, exampleUser) {
				// 		insertUser(t, db, exampleUser)
				// 	}

				// 	insertProduct(t, exampleProducts[0])
				// }
			} else {
				handlers.Storage = database.New(nil)
			}

			// Now the request
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.URL.RawQuery = ts.QueryParams.Encode()

			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, ts.ExpectedStatusCode, rec.Code)
			assertResponseError(t, ts.ExpectedError, decodeResponseMessage(t, rec))
		})
	}
}
