package main

import (
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Action map[string]string
type Actions []Action

type Scenario struct {
	Name    string  `name`
	Actions Actions `actions`
}

type Scenarios []Scenario

type Playscript struct {
	FullPath   string
	BaseName   string
	BaseDir    string
	Scenarios  Scenarios `scenarios`
	Before     Actions   `before`
	After      Actions   `after`
	BeforeEach Actions   `before_each`
	AfterEach  Actions   `after_each`
}

type ScenarioFile struct {
	FullPath  string
	BaseName  string
	BaseDir   string
	Scenarios Scenarios
}

func NewPlayscript(inputFilePath string) (*Playscript, error) {
	fullPath, err := filepath.Abs(inputFilePath)

	if err != nil {
		return nil, err
	}

	playscript := &Playscript{}

	playscript.FullPath = fullPath
	playscript.BaseDir = filepath.Dir(fullPath)
	playscript.BaseName = filepath.Base(fullPath)

	data, err := ReadFile(fullPath)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &playscript)

	if err != nil {
		return nil, err
	}

	return playscript, nil
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
	data, err := ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &scenarios)

	if err != nil {
		log.Fatal(err)
	}

	return scenarios, nil
}
