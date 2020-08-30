package main

import (
	"log"

	"github.com/coffemanfp/beppin-server/database"
	"github.com/coffemanfp/beppin-server/utils"
)

func main() {
	storage, err := database.Get()
	if err != nil {
		log.Fatalln(err)
	}

	db := storage.GetDB()

	schemaBytes, err := utils.GetFilebytes(schemaFile)
	if err != nil {
		log.Fatalln("failed to read the schema file: ", err)
	}

	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		log.Fatalln("failed to execute the schema: ", err)
	}

	log.Println("Schema executed successfully!!")

	if !withExamples {
		return
	}

	examplesBytes, err := utils.GetFilebytes(examplesFile)
	if err != nil {
		log.Fatalln("failed to read the examples file: ", err)
	}

	_, err = db.Exec(string(examplesBytes))
	if err != nil {
		log.Fatalln("failed to execute the examples: ", err)
	}

	log.Println("Examples executed successfully!!")
}

func init() {
	initFlags()
	initSettings()
}
