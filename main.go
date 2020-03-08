package main

import (
	"github.com/joho/godotenv"
    "log"
    "os"
	"errors"
	"path/filepath"
)

func main() {

}


func doesFileExist(name string){
	if _, err := os.Stat(""); os.IsNotExist(err) {
	// path/to/whatever does not exist
	}
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
            log.Fatal(err)
	}
	
    return dir
}