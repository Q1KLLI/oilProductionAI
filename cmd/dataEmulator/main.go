package main

import (
	"github.com/Q1KLLI/oilProductionAI/internal/dbController"
	"github.com/Q1KLLI/oilProductionAI/internal/emulatorConfig"
	"log"
)

const ConfigFilename = `C:\repositories\oilProductionAI\config\emulatorConfig.yml`

func main() {

	conf := emulatorConfig.NewConfig(ConfigFilename)

	db, err := dbController.NewDatabase(conf.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	wells, err := db.Wells()
	if err != nil {
		log.Fatal(err)
	}

}
