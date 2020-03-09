package main

import (
	"github.com/joho/godotenv"
    "log"
    "os"
	"path/filepath"
)

func main() {
  file := ".env"

  if FileExist(file) {
		err := godotenv.Load(file)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}


func FileExist(name string) bool{
	dir, _ := CurrentDir()
	if _, err := os.Stat(dir + name); os.IsNotExist(err) {
		return true
	}
	return false
}

func CurrentDir() (dir string, err error){
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
            log.Fatal(err)
	}
	
    return
}