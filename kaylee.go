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

func init() {
    flag.StringVar(&configString, "c", "not_set", "JSON string containing config")
    flag.StringVar(&configString, "config", "not_set", "JSON string containing config")

    flag.BoolVar(&printExample, "e", false, "print example json")
    flag.BoolVar(&printExample, "example", false, "print example json")

    flag.BoolVar(&printUsage, "h", false, "print usage")
    flag.BoolVar(&printUsage, "help", false, "print usage")

    flag.BoolVar(&verbose, "v", false, "enable verbose output")
    flag.BoolVar(&verbose, "verbose", false, "enable verbose output")
}

func main() {

	flag.Parse()

	if printExample { PrintExample() }

	if printUsage { PrintUsage() }

    if configString == "not_set" {
        LogError("config [-c|-config] flag is required")
        PrintUsage()
    }

    LogVerbose("verbose mode enabled")

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

Usage:`

	fmt.Println(helpText)
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
    }
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

func LogError(err string, args... interface{}) {
    if len(args) > 0 {
        err = fmt.Sprintf(err, args...)
    }
	fmt.Fprintf(os.Stderr, "\033[01;31mKAYLE_ERROR: %s\033[00;00m\n", err)
	os.Exit(1)
}

func LogVerbose(line string, args... interface{}) {
	if verbose == true {
        if len(args) > 0 {
            line = fmt.Sprintf(line, args...)
        }
		fmt.Printf("verbose: %s\n", line)
	}
}