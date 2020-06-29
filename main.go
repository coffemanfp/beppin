package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
	"github.com/coffemanfp/beppin-server/router"
	"github.com/labstack/echo"
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

	router.NewRouter(e)
	err = config.NewLogger(e, "logs/server.log")
	if err != nil {
		log.Fatalf("failed to set logger:\n%s", err)
	}

	log.Println(e.Start(fmt.Sprintf(":%d", settings.Port)))
}

func init() {
	initSettings()
	initDatabase()
}

func initSettings() {
	err := config.SetSettingsByFile(configFileDef)
	if err != nil {
		log.Fatalln(err)
	}
}

func initDatabase() {
	_, err := database.OpenConn()
	if err != nil {
		log.Fatalln(err)
	}
}

func initFlags() {
	flag.StringVar(&configFile, "config-file", configFileDef, "Config file for the server settings.")

	flag.Parse()
}
