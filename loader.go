package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadSenarios(filePath string) ([]Scenario, error) {
	fileName := filePath
	scenarios := []Scenario{}
	reader, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	Debug("Load %s", fileName)

	defer reader.Close()

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &scenarios)

	if err != nil {
		log.Fatal(err)
	}

	return scenarios, nil
}
