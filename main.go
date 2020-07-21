package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	settings, err := config.GetSettings()
	if err != nil {
		log.Fatalf("failed to get settings:\n%s", err)
	}

	e := echo.New()

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

	// Create routes
	err = router.NewRouter(e)
	if err != nil {
		log.Fatalf("failed to set router:\n%s", err)
	}

	// Config logger
	err = config.NewLogger(e, settings.LogsFile)
	if err != nil {
		log.Fatalf("failed to set logger:\n%s", err)
	}

	// Run server and print if fails.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", settings.Port)))
}

func init() {
	initSettings()
	initDatabase()
}
