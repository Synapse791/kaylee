package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configString    string
	verbose         bool
)

func main() {

	flag.StringVar(&configString, "c", "not_set", "JSON string containing config")
	flag.BoolVar(&verbose, "v", false, "enable verbose output")
	flag.Parse()

	config, err := GetConfig(configString)
	if err != nil {
		LogError(err.Error())
	}

	err = FindReplace(config)
	if err != nil {
		LogError(err.Error())
	}

	os.Exit(0)

}

func LogError(err string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	os.Exit(1)
}

func LogVerbose(line string) {
	if verbose == true {
		fmt.Printf("verbose: %s\n", line)
	}
}