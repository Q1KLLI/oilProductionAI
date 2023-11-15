package emulatorConfig

import (
	"github.com/Q1KLLI/oilProductionAI/internal/dbController"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Configuration struct {
	DBConfig dbController.Config `yaml:"DBConfig"`
}

func NewConfig(filename string) Configuration {

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	c := Configuration{}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalln(err)
	}

	return c
}
