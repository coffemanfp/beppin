package main

import (
	"flag"
	"log"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
)

func initSettings() {
	err := config.SetMigrationsSettingsByEnv()
	if err != nil {
		log.Fatalln("failed to configure env settings: ", err)
	}

	if configFile != "" {
		err = config.SetMigrationsSettingsByFile(configFile)
		if err != nil {
			log.Fatalln("failed to configure settings: ", err)
		}
	}
}

func initDatabase() {
	_, err := database.OpenConn()
	if err != nil {
		log.Fatalln("failed to start the database: ", err)
	}
}

func initFlags() {
	flag.BoolVar(&withExamples, "with-examples", false, "Add examples to the database.")
	flag.StringVar(&configFile, "config-file", "migrations/config.yaml", "Config file for the database settings.")
	flag.StringVar(&schemaFile, "schema-file", "migrations/schema.sql", "Schema to execute")
	flag.StringVar(&examplesFile, "examples-file", "migrations/examples.sql", "Examples to execute")

	flag.Parse()
}
