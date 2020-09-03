package handlers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/coffemanfp/beppin/config"
	"github.com/coffemanfp/beppin/database"
	dbm "github.com/coffemanfp/beppin/database/models"
	"github.com/coffemanfp/beppin/handlers"
	"github.com/coffemanfp/beppin/helpers"
	"github.com/coffemanfp/beppin/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func init() {
	config.SetDefaultSettings()
	txdb.Register("beppin_test", "postgres", config.GlobalSettings.Database.URL)
}

var token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJsYW5ndWFnZSI6ImVzLUVTIiwidXNlcm5hbWUiOiJjb2ZmZW1hbmZwIiwidGhlbWUiOiJsaWdodCJ9fQ.GJcykxeN4yfE7CVi1xu4zVYstPgODCuNtrgq4T11gA4"

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

	storage := database.New(db)
	assert.NotNil(t, storage)

	handlers.Storage = storage
}

func insertUser(t *testing.T, db *sql.DB, user models.User) {
	t.Helper()

	userBytes, err := yaml.Marshal(user)
	assert.Nil(t, err)

	p := polluter.New(polluter.PostgresEngine(db))

	assert.Nil(t, p.Pollute(bytes.NewReader(userBytes)))
	return
}

func existsUser(t *testing.T, user models.User) (exists bool) {
	t.Helper()

	userDB, err := helpers.ParseModelToDBModel(user)
	assert.Nil(t, err)

	exists, err = handlers.Storage.ExistsUser(userDB.(dbm.User))
	assert.Nil(t, err)
	return
}

func insertLanguage(t *testing.T, language models.Language) (id int) {
	t.Helper()

	languageDB, err := helpers.ParseModelToDBModel(language)
	assert.Nil(t, err)

	id, err = handlers.Storage.CreateLanguage(languageDB.(dbm.Language))
	assert.Nil(t, err)
	return
}

func existsLanguage(t *testing.T, language models.Language) (exists bool) {
	t.Helper()

	languageDB, err := helpers.ParseModelToDBModel(language)
	assert.Nil(t, err)

	exists, err = handlers.Storage.ExistsLanguage(languageDB.(dbm.Language))
	assert.Nil(t, err)
	return
}

func insertProduct(t *testing.T, product models.Product) (id int) {
	t.Helper()

	productDB, err := helpers.ParseModelToDBModel(product)
	assert.Nil(t, err)

	id, err = handlers.Storage.CreateProduct(productDB.(dbm.Product))
	assert.Nil(t, err)
	return
}

func getProduct(t *testing.T, id int) (product models.Product, err error) {
	t.Helper()

	productDB, err := handlers.Storage.GetProduct(dbm.Product{ID: int64(id)})
	assert.Nil(t, err)

	product = helpers.ShouldParseDBModelToModel(productDB).(models.Product)
	return
}

func restartSequences(t *testing.T, db *sql.DB) {
	t.Helper()

	query := `
	SELECT
		'ALTER SEQUENCE "' || c.relname || '" RESTART'
	FROM
		pg_class c
	WHERE
		c.relkind = 'S'
	`

	stmt, err := db.Prepare(query)
	assert.Nil(t, err)
	defer stmt.Close()

	rows, err := stmt.Query()
	assert.Nil(t, err)

	var queryRestartSequence string

	for rows.Next() {
		err = rows.Scan(&queryRestartSequence)
		assert.Nil(t, err)

		stmtRestartSequence, err := db.Prepare(queryRestartSequence)
		assert.Nil(t, err)
		defer stmtRestartSequence.Close()

		_, err = stmtRestartSequence.Exec()
		assert.Nil(t, err)
	}
}
