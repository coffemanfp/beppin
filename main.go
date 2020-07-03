package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
	"github.com/coffemanfp/beppin-server/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	configFile    string
	configFileDef string = "config.yaml"
)

func main() {
	settings, err := config.GetSettings()
	if err != nil {
		log.Fatalf("failed to get settings:\n%s", err)
	}

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	// CORS
	e.Use(middleware.CORS())
	// Create routes
	router.NewRouter(e)

	// Config logger
	err = config.NewLogger(e, "logs/server.log")
	if err != nil {
		log.Fatalf("failed to set logger:\n%s", err)
	}

	// Run server and print if fails.
	log.Println(e.Start(fmt.Sprintf(":%d", settings.Port)))
}

func init() {
	initSettings()
	initDatabase()
}

func initSettings() {
	err := config.SetSettingsByFile(configFileDef)
	if err != nil {
		log.Fatalln("failed to configure settings:\n", err)
	}

	err = config.SetSettingsByEnv()
	if err != nil {
		log.Fatalln("failed to configure env settings:\n", err)
	}
}

func initDatabase() {
	_, err := database.OpenConn()
	if err != nil {
		log.Fatalln("failed to start the database:\n", err)
	}
}

func initFlags() {
	flag.StringVar(&configFile, "config-file", configFileDef, "Config file for the server settings.")

	flag.Parse()
}
