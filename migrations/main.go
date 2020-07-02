package main

import (
	"flag"
	"log"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/database"
	"github.com/coffemanfp/beppin-server/utils"
)

var (
	withExamples    bool
	withExamplesDef bool
	configFile      string
	configFileDef   string = "config.yaml"
	schemaFile      string
	schemaFileDef   string = "migrations/schema.sql"
	examplesFile    string
	examplesFileDef string = "migrations/examples.sql"
)

func main() {
	err := config.SetSettingsByFile(configFile)
	if err != nil {
		log.Fatalln("failed to configure settings:\n", err)
	}

	db, err := database.Get()
	if err != nil {
		log.Fatalln(err)
	}

	schemaBytes, err := utils.GetFilebytes(schemaFile)
	if err != nil {
		log.Fatalln("failed to read the schema file:\n", err)
	}

	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		log.Fatalln("failed to execute the schema:\n", err)
	}

	log.Println("Schema executed successfully!!")

	if !withExamples {
		return
	}

	examplesBytes, err := utils.GetFilebytes(examplesFile)
	if err != nil {
		log.Fatalln("failed to read the examples file:\n", err)
	}

	_, err = db.Exec(string(examplesBytes))
	if err != nil {
		log.Fatalln("failed to execute the examples:\n", err)
	}

	log.Println("Examples executed successfully!!")
}

func init() {
	flag.BoolVar(&withExamples, "with-examples", withExamplesDef, "Add examples to the database.")
	flag.StringVar(&configFile, "config-file", configFileDef, "Config file for the database settings.")
	flag.StringVar(&schemaFile, "schema-file", schemaFileDef, "Schema to execute")
	flag.StringVar(&examplesFile, "examples-file", examplesFileDef, "Examples to execute")

	flag.Parse()
}
