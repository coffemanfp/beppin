package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.Ping(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "pong", rec.Body.String())
	}
}
