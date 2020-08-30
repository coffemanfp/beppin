package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
	dbm "github.com/coffemanfp/beppin-server/database/models"
	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/helpers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/testify/assert"
)

func init() {
	exampleTime = time.Now()
	config.SetDefaultSettings()
}

var token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJsYW5ndWFnZSI6ImVzLUVTIiwidXNlcm5hbWUiOiJjb2ZmZW1hbmZwIiwidGhlbWUiOiJsaWdodCJ9fQ.GJcykxeN4yfE7CVi1xu4zVYstPgODCuNtrgq4T11gA4"
var exampleTime time.Time

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
		Name:        fmt.Sprintf("Product at %d", exampleTime.UnixNano()),
		Description: "Product description",
		Categories: []string{
			"Food",
		},
	},
	models.Product{
		UserID:      1,
		Name:        fmt.Sprintf("Product at %d", exampleTime.UnixNano()),
		Description: "Product description 2",
		Categories: []string{
			"Tech",
		},
	},
	models.Product{
		UserID:      1,
		Name:        fmt.Sprintf("Product at %d", exampleTime.UnixNano()),
		Description: "Product description 3",
		Categories: []string{
			"Clothes",
		},
	},
}

// Server helpers

func assertResponseError(t *testing.T, expectedError string, m models.ResponseMessage) {
	t.Helper()

	if assert.NotNil(t, m) {
		assert.Equal(t, expectedError, m.Error)
	}
}

func assertResponseMessage(t *testing.T, expectedMessage string, m models.ResponseMessage) {
	t.Helper()

	if assert.NotNil(t, m) {
		assert.Equal(t, expectedMessage, m.Message)
	}
}

func decodeResponseMessage(t *testing.T, rec *httptest.ResponseRecorder) (m models.ResponseMessage) {
	t.Helper()

	assert.Nil(t, json.NewDecoder(rec.Body).Decode(&m))
	return
}

func setJWTMiddleware(t *testing.T, e *echo.Echo) {
	t.Helper()

	jwtConfig := middleware.JWTConfig{
		Claims:      &models.Claim{},
		SigningKey:  []byte(config.GlobalSettings.SecretKey),
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
