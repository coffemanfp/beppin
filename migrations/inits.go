package main

import (
	"flag"
)

func initFlags() {
	flag.BoolVar(&withExamples, "with-examples", false, "Add examples to the database.")
	flag.StringVar(&configFile, "config-file", "migrations/config.yaml", "Config file for the database settings.")
	flag.StringVar(&schemaFile, "schema-file", "migrations/schema.sql", "Schema to execute")
	flag.StringVar(&examplesFile, "examples-file", "migrations/examples.sql", "Examples to execute")

	flag.Parse()
}
