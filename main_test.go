package main

import (
	"testing"
	"path/filepath"
	"os"
)


//Should return true if the name of the file exist in the current directory
func TestFileExist(t *testing.T){
	received := FileExist(".env")
	expected := true

	if received != expected {
       t.Error("Expected: ", expected, "Received: ", received)
	}
}

//Should return the directory of the where the file is being run on the computer
func TestCurrentDir(t *testing.T) {

	received, err := filepath.Abs(filepath.Dir(os.Args[0]))
	expected, err := CurrentDir()

	if expected != received {
       t.Error("Expected: ", expected, "Received: ", received)
	}
	
	if err != nil {
		t.Error("Error: ", err)
	}
}
