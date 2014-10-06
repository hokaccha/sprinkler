package main

import (
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

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
