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
	configFile    string
	configFileDef string = "config.yaml"
)

var filesToUpload = make(chan *os.File)

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
