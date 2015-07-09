package main

import (
    "encoding/json"
    "fmt"
)

type KayleeConfig struct {
    Files []FileConfig
}

type FileConfig struct {
    Path     string            `json:"path"`
    Patterns map[string]string `json:"patterns"`
}

func GetConfig(conf string) (KayleeConfig, error) {
    var config KayleeConfig
    var files []FileConfig

    err := json.Unmarshal([]byte(conf), &files)
    if err != nil {
        LogVerbose(err.Error())
        return config, fmt.Errorf("invalid JSON")
    }

    config.Files = files

    return config, nil
}