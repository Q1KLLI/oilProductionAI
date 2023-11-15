package main

import (
	"fmt"
	"github.com/Q1KLLI/oilProductionAI/internal/dbController"
	"github.com/Q1KLLI/oilProductionAI/internal/emulatorConfig"
	"log"
)

func main() {

	conf := emulatorConfig.NewConfig(`D:\repo\oilProductionAI\config\emulatorConfig.yml`)

	db, err := dbController.NewDatabase(conf.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	wellIDs, err := db.WellIDs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(wellIDs)

}
