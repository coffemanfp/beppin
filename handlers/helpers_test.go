package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	errs "github.com/coffemanfp/beppin-server/errors"
	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

var token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJsYW5ndWFnZSI6ImVzLUVTIiwidXNlcm5hbWUiOiJjb2ZmZW1hbmZwIiwidGhlbWUiOiJsaWdodCJ9fQ.GJcykxeN4yfE7CVi1xu4zVYstPgODCuNtrgq4T11gA4"
var exampleTime = time.Now()

var exampleLanguage = models.Language{
	Code:   "es-ES",
	Status: "available",
}

var exampleUser = models.User{
	Language: "es-ES",
	Name:     "Franklin",
	Username: "coffemanfp",
	Password: "24061502",
	LastName: "Peñaranda",
	Theme:    "dark",
	Birthday: &exampleTime,
}

var exampleProducts = models.Products{
	models.Product{
		UserID:      1,
		Name:        "Product name",
		Description: "Product description",
		Categories: []string{
			"Food",
		},
	},
	models.Product{
		UserID:      1,
		Name:        "Product name 2",
		Description: "Product description 2",
		Categories: []string{
			"Tech",
		},
	},
	models.Product{
		UserID:      1,
		Name:        "Product name 3",
		Description: "Product description 3",
		Categories: []string{
			"Clothes",
		},
	},
}

// Server helpers

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
	t.Helper()

	assert.Equal(t, http.StatusCreated, rec.Code)

	var m models.ResponseMessage
	assert.Nil(t, json.NewDecoder(rec.Body).Decode(&m))
	assert.Equal(t, "Created.", m.Message)
}

func setJWTMiddleware(t *testing.T, e *echo.Echo) {
	t.Helper()

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

// Database helpers

func setStorage(t *testing.T) {
	t.Helper()

	storage, err := database.NewDefault()
	assert.Nil(t, err)
	assert.NotNil(t, storage)
	handlers.Storage = storage
}

func insertUser(t *testing.T, user models.User) {
	t.Helper()

	userDB, err := helpers.ParseModelToDBModel(user)
	assert.Nil(t, err)

	assert.Nil(t, handlers.Storage.CreateUser(userDB.(dbm.User)))
}

func existsUser(t *testing.T, user models.User) (exists bool) {
	t.Helper()

	userDB, err := helpers.ParseModelToDBModel(user)
	assert.Nil(t, err)

	exists, err = handlers.Storage.ExistsUser(userDB.(dbm.User))
	assert.Nil(t, err)
	return
}

func insertLanguage(t *testing.T, language models.Language) {
	t.Helper()

	languageDB, err := helpers.ParseModelToDBModel(language)
	assert.Nil(t, err)

	assert.Nil(t, handlers.Storage.CreateLanguage(languageDB.(dbm.Language)))
}

func existsLanguage(t *testing.T, language models.Language) (exists bool) {
	t.Helper()

	languageDB, err := helpers.ParseModelToDBModel(language)
	assert.Nil(t, err)

	exists, err = handlers.Storage.ExistsLanguage(languageDB.(dbm.Language))
	assert.Nil(t, err)
	return
}

func insertProduct(t *testing.T, product models.Product) {
	t.Helper()

	productDB, err := helpers.ParseModelToDBModel(product)
	assert.Nil(t, err)

	assert.Nil(t, handlers.Storage.CreateProduct(productDB.(dbm.Product)))
}
