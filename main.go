package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coffemanfp/beppin/config"
	"github.com/coffemanfp/beppin/database"
	"github.com/coffemanfp/beppin/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stretchr/gomniauth"
)

func main() {
	e := echo.New()

	gomniauth.SetSecurityKey(config.GlobalSettings.SecretKey)

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

	// Pass the database connection to the handlers
	handlers.Storage, _ = database.Get()

	// Config logger
	err := config.NewLogger(e, config.GlobalSettings.LogsFile)
	if err != nil {
		log.Fatalf("failed to set logger: %v", err)
	}

	// Run server and print if fails.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.GlobalSettings.Port)))
}

func init() {
	initFlags()
	initSettings()
	initDatabase()
}
