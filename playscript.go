package main

import (
	"path/filepath"

	"github.com/hokaccha/sprinkler/utils"
)

type Actions []map[string]interface{}

type Scenario struct {
	Name    string      `name`
	Actions Actions     `actions`
	Include string      `include`
	Tags    interface{} `tags`
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

func NewPlayscript(inputFilePath string) (*Playscript, error) {
	fullPath, err := filepath.Abs(inputFilePath)

	if err != nil {
		return nil, err
	}

	playscript := &Playscript{}

	playscript.FullPath = fullPath
	playscript.BaseDir = filepath.Dir(fullPath)
	playscript.BaseName = filepath.Base(fullPath)

	err = utils.LoadYAML(fullPath, &playscript)

	if err != nil {
		return nil, err
	}

	return playscript, nil
}
