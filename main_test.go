package main

import (
	"testing"
	"fmt"
)

func TestEnvParse(t *testing.T) {
    result, error := EnvParse()

	fmt.Println(result)

	if error != nil {
       t.Errorf("Error not expected from TestEnvParse")
    }
}


func TestGetCurrentDir(t *testing.T) {
	result, error := 
}