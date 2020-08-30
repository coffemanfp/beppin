package main

import (
	"flag"
	"log"

	"github.com/coffemanfp/beppin-server/config"
)

var (
	withExamples   bool
	readConfigFile bool
	readEnvVars    bool
	schemaFile     string
	examplesFile   string
)

func initFlags() {
	flag.BoolVar(&readConfigFile, "read-config-file", false, "Checks if read the config file or not.")
	flag.BoolVar(&readEnvVars, "read-env-vars", false, "Checks if read the environment vars or not.")

	flag.StringVar(&schemaFile, "schema-file", "migrations/schema.sql", "Schema to execute")
	flag.BoolVar(&withExamples, "with-examples", false, "Add examples to the database.")
	flag.StringVar(&examplesFile, "examples-file", "migrations/examples.sql", "Examples to execute")

	flag.Parse()
}

func initSettings() {
	err := config.SetDefaultSettings()
	if err != nil {
		log.Fatalln("failed to configure default settings:", err)
	}

	if readEnvVars {
		err = config.SetSettingsByEnv()
		if err != nil {
			log.Fatalln("failed to configure environment vars settings:", err)
		}
	}

	if readConfigFile {
		err := config.SetSettingsByFile()
		if err != nil {
			log.Fatalln("failed to configure file settings: ", err)
		}
	}
}
