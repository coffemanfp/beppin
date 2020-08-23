package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping is a simple handler ping that responses pong
func Ping(c echo.Context) (err error) {
	return c.String(http.StatusOK, "pong")
}
