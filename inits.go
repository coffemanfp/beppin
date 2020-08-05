package main

import (
	"flag"
	"log"
	"os"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
)

// Flags
var (
	configFile string
)

var filesToUpload = make(chan *os.File)

func initSettings() {
	config.SetDefaultSettings()

	err := config.SetSettingsByEnv()
	if err != nil {
		log.Fatalln("failed to configure env settings:\n", err)
	}

	if configFile != "" {
		err := config.SetSettingsByFile(configFile)
		if err != nil {
			log.Fatalln("failed to configure file settings:\n", err)
		}
	}
}

func initDatabase() {
	_, err := database.OpenConn()
	if err != nil {
		log.Fatalln("failed to start the database:\n", err)
	}
}

func initFlags() {
	flag.StringVar(&configFile, "config-file", "config.yaml", "Config file for the server settings.")

	flag.Parse()
}
