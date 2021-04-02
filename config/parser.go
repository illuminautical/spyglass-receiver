package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func readFile(path string) []byte {

	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error("Failed to read yaml file", err)
	}

	return file
}

func (c *ApplicationConfig) LoadConfig(path string) *ApplicationConfig {

	file := readFile(path)

	err := yaml.Unmarshal(file, c)
	if err != nil {
		log.Error("Failed to parse yaml", err)
	}

	return c
}