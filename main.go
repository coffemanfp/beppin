package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/gomniauth"
)

func main() {
	settings := config.GetSettings()

	e := echo.New()

	gomniauth.SetSecurityKey(settings.SecretKey)

	// Middlewares
	{
		e.Use(
			// middleware.Logger(),
			middleware.Recover(),
			middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{
					http.MethodPost,
					http.MethodGet,
					http.MethodPut,
					http.MethodDelete,
				},
				Skipper: middleware.DefaultSkipper,
			}),
		)
	}

	// Remove Trailing URL Slash
	e.Pre(middleware.RemoveTrailingSlash())
	e.Static("static", "assets")

	// Create routes
	newRouter(e)

	// Config logger
	err := config.NewLogger(e, settings.LogsFile)
	if err != nil {
		log.Fatalf("failed to set logger: %v", err)
	}

	// Run server and print if fails.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", settings.Port)))
}

func init() {
	initFlags()
	initSettings()
	initDatabase()
}
