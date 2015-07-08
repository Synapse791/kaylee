package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	configString    string
	printExample    bool
	printUsage      bool
	verbose         bool
)

func main() {

	flag.StringVar(&configString, "config", "not_set", "JSON string containing config")
	flag.BoolVar(&printExample, "example", false, "print example json")
	flag.BoolVar(&printUsage, "help", false, "print usage")
	flag.BoolVar(&verbose, "v", false, "enable verbose output")
	flag.Parse()

	if printExample { PrintExample() }

	if printUsage || configString == "not_set" { PrintUsage() }

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

func PrintUsage() {
	helpText := `Kaylee - find and replace inside files using json config

  Pass in a JSON string to specify a group of files. Each of these files has a group of find and replace values.
  Kaylee runs through these groups and replaces the find strings with the replace strings, then moves onto the next file.

`

	fmt.Print(helpText)
	flag.PrintDefaults()
	os.Exit(0)
}

func PrintExample() {
	example := `[
  {
    "path"     : "/tmp/file1",
    "patterns" : {
        "replace_me" : "with_me",
        "password"   : "letmein"
      },
  },
  {
    "path"     : "/tmp/file2",
    "patterns" : {
        "FINDME1" : "put me in the file"
    }
  }
]`

	fmt.Print(example)
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