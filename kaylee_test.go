package main

import (
	"testing"
	"fmt"
	"strings"
	"io/ioutil"
	"syscall"
)

func TestGetConfigWithValidJson(t *testing.T) {
	testConfig := `[
	  {
	    "path"     : "/tmp/file",
	    "patterns" : [
	      {
	        "find"    : "1",
	        "replace" : "2"
	      }
	    ]
	  }
	]`

	_, err := GetConfig(testConfig)

	if err != nil {
		t.Errorf("expected success :: got error '%s'", err.Error())
	}
}

func TestGetConfigWithInvalidJson(t *testing.T) {
	testConfig := `I AM NOT JSON`

	_, err := GetConfig(testConfig)

	if err == nil {
		t.Error("expected error 'invalid JSON' :: got nil")
	}

	if err != nil && strings.Contains(err.Error(), "invalid JSON") != true {
		t.Errorf("expected error 'invalid JSON' :: got error '%s'", err.Error())
	}
}

func TestFindReplaceWithNonexistentFile(t *testing.T) {
	testConfig := `[
	  {
	    "path"     : "not_exist",
	    "patterns" : [
	      {
	        "find"    : "was",
	        "replace" : "am"
	      }
	    ]
	  }
	]`

	config, confErr := GetConfig(testConfig)
	if confErr != nil {
		t.Errorf("expected to be able to get config :: got error '%s'", confErr.Error())
	}

	fiReErr := FindReplace(config)
	if fiReErr == nil {
		t.Error("expected error 'file not_exist does not exist' :: got nil")
	}

	if fiReErr != nil && strings.Contains(fiReErr.Error(), "file not_exist does not exist") != true {
		t.Errorf("expected error 'file not_exist does not exist' :: got '%s'", fiReErr.Error())
	}

}

func TestFindReplaceWithUnreadableFile(t *testing.T) {
	f, err := ioutil.TempFile("", "file1")
	if err != nil { t.Fatal("failed to create test file") }
	defer syscall.Unlink(f.Name())
	ioutil.WriteFile(f.Name(), []byte("content"), 0)

	testConfig := `[
	  {
	    "path"     : "%s",
	    "patterns" : [
	      {
	        "find"    : "was",
	        "replace" : "am"
	      }
	    ]
	  }
	]`

	config, confErr := GetConfig(fmt.Sprintf(testConfig, f.Name()))
	if confErr != nil {
		t.Errorf("expected to be able to get config :: got error '%s'", confErr.Error())
	}

	f.Chmod(0200)

	fiReErr := FindReplace(config)

	if fiReErr == nil {
		t.Errorf("expected error 'failed to read file %s' :: got nil", f.Name())
	}

	if fiReErr != nil && fiReErr.Error() != fmt.Sprintf("failed to read file %s", f.Name()) {
		t.Errorf("expected error 'failed to read file %s' :: got error '%s'", f.Name(), fiReErr.Error())
	}
}

func TestFindReplaceForSuccess(t *testing.T) {

	expected := "i am groot"

	f, err := ioutil.TempFile("", "file1")
	if err != nil { t.Fatal("failed to create test file") }
	defer syscall.Unlink(f.Name())
	ioutil.WriteFile(f.Name(), []byte(expected), 0)

	testConfig := `[
	  {
	    "path"     : "%s",
	    "patterns" : [
	      {
	        "find"    : "was",
	        "replace" : "am"
	      }
	    ]
	  }
	]`

	config, confErr := GetConfig(fmt.Sprintf(testConfig, f.Name()))
	if confErr != nil {
		t.Errorf("expected to be able to get config :: got error '%s'", confErr.Error())
	}

	fiReErr := FindReplace(config)
	if fiReErr != nil {
		t.Errorf("expected to succesffully configure file :: got error '%s'", fiReErr.Error())
	}

	recieved, readErr := ioutil.ReadFile(f.Name())
	if readErr != nil {
		t.Errorf("failed to reaed file %s", f.Name())
	}

	recievedString := string(recieved)

	if recievedString != expected {
		t.Errorf("expected '%s' :: got '%s'", expected, recievedString)
	}

}










