package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func FindReplace(config KayleeConfig) error {

	for _, file := range config.Files {
		if _, err := os.Stat(file.Path); os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", file.Path)
		}
	}

	for _, file := range config.Files {
		LogVerbose(fmt.Sprintf("configuring %s", file.Path))

		data, err := ioutil.ReadFile(file.Path)
		if err != nil {
			return fmt.Errorf("failed to read file %s", file.Path)
		}

		newContents := string(data)

		for _, p := range file.Patterns {
			LogVerbose(fmt.Sprintf("replacing '%s' with '%s' in %s", p.Find, p.Replace, file.Path))
			newContents = strings.Replace(newContents, p.Find, p.Replace, -1)
		}

		LogVerbose(fmt.Sprintf("writing file %s", file.Path))
		err = ioutil.WriteFile(file.Path, []byte(newContents), 0)
		if err != nil {
			return fmt.Errorf("failed to write file %s", file.Path)
		}

		LogVerbose(fmt.Sprintf("successfully configured %s", file.Path))

	}

	return nil

}