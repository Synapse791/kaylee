package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func FindReplace(config KayleeConfig) error {

    for _, file := range config.Files {
        if _, err := os.Stat(file.Path); os.IsNotExist(err) {
            LogVerbose(err.Error())
            return fmt.Errorf("file %s does not exist", file.Path)
        }
    }

    for _, file := range config.Files {
        LogVerbose("configuring %s", file.Path)

        data, err := ioutil.ReadFile(file.Path)
        if err != nil {
            LogVerbose(err.Error())
            return fmt.Errorf("failed to read file %s", file.Path)
        }

        newContents := string(data)

        for find, replace := range file.Patterns {
            if strings.Contains(newContents, find) {
                LogVerbose("replacing '%s' with '%s' in %s", find, replace, file.Path)
                newContents = strings.Replace(newContents, find, replace, -1)
            } else {
                return fmt.Errorf("'%s' not found in '%s'", find, file.Path)
            }

        }

        LogVerbose("writing file %s", file.Path)
        err = ioutil.WriteFile(file.Path, []byte(newContents), 0)
        if err != nil {
            LogVerbose(err.Error())
            return fmt.Errorf("failed to write file %s", file.Path)
        }

        LogVerbose("successfully configured %s", file.Path)

    }

    return nil

}