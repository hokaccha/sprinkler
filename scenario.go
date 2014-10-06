package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Action map[string]string

type Scenario struct {
	Name    string   `name`
	Actions []Action `actions`
}

type Scenarios []Scenario

type ScenarioFile struct {
	FullPath  string
	BaseName  string
	BaseDir   string
	Scenarios Scenarios
}

func NewSenarioFile(inputFilePath string) (*ScenarioFile, error) {
	fullPath, err := filepath.Abs(inputFilePath)

	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(fullPath)
	baseName := filepath.Base(fullPath)
	scenarios, err := LoadSenarios(fullPath)

	if err != nil {
		return nil, err
	}

	return &ScenarioFile{
		FullPath:  fullPath,
		BaseName:  baseName,
		BaseDir:   baseDir,
		Scenarios: scenarios,
	}, nil
}

func LoadSenarios(filePath string) (Scenarios, error) {
	scenarios := Scenarios{}
	reader, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	Debug("Load %s", filePath)

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
