package main

import (
	"fmt"
	"encoding/json"
)

type KayleeConfig struct {
	Files   []FileConfig
}

type FileConfig struct {
	Path        string          `json:"path"`
	Patterns    []PatternSet    `json:"patterns"`
}

type PatternSet struct {
	Find    string  `json:"find"`
	Replace string  `json:"replace"`
}

func GetConfig(conf string) (KayleeConfig, error) {
	var config KayleeConfig
	var files []FileConfig

	err := json.Unmarshal([]byte(conf), &files)
	if err != nil {
		return config, fmt.Errorf("invalid JSON")
	}

	config.Files = files

	return config, nil
}